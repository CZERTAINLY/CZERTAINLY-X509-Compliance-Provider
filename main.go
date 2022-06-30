package main

import (
	"CZERTAINLY-X509-Compliance-Provider/cmd"
	"CZERTAINLY-X509-Compliance-Provider/cmd/compliance"
	"CZERTAINLY-X509-Compliance-Provider/cmd/info"
	"CZERTAINLY-X509-Compliance-Provider/cmd/rules"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port string `yaml:"port", envconfig:"SERVER_PORT"`
	} `yaml:"server"`
	Log struct {
		Level string `yaml:"level", envconfig:"LOG_LEVEL"`
	} `yaml:"log"`
}

var CompositeRouter = mux.NewRouter()
var config Config

func init() {
	readConfigFile(&config)
	readEnvironmentVariables(&config)
}

func main() {

	var httpAddr = flag.String("http", ":"+config.Server.Port, "http listen address")
	rawJSON := []byte(`{
		"level": "` + strings.ToLower(config.Log.Level) + `",
		"encoding": "json",
		"outputPaths": ["stdout"],
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {
		  "messageKey": "message",
		  "levelKey": "level",
		  "levelEncoder": "lowercase",
		  "timeKey": "timestamp",
		  "timeEncoder": "rfc3339"
		}
	  }`)
	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	logger.Info("Starting CZERTAINLY-X509-Compliance-Provider")
	logger.WithOptions()
	defer logger.Sync()
	sugar := logger.Sugar()

	flag.Parse()
	ctx := context.Background()
	errs := make(chan error)

	infoSrv, complianceSrv, rulesSrv := createService(sugar)
	infoEndpoints, complianceEndpoints, rulesEndpoints := createEndPoints(infoSrv, complianceSrv, rulesSrv)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := cmd.NewHttpServer(CompositeRouter, ctx, infoEndpoints, complianceEndpoints, rulesEndpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()
	fmt.Print(<-errs)
}

func createService(logger *zap.SugaredLogger) (info.Service, compliance.Service, rules.Service) {
	var infoSrv info.Service
	{
		infoSrv = info.NewService(logger)
	}

	var complianceSrv compliance.Service
	{
		complianceSrv = compliance.NewService(logger)
	}

	var rulesSrv rules.Service
	{
		rulesSrv = rules.NewService(logger)
	}
	return infoSrv, complianceSrv, rulesSrv
}

func createEndPoints(infoService info.Service, complianceService compliance.Service, rulesService rules.Service) (info.EndPoints, compliance.EndPoints, rules.EndPoints) {
	infoEndpoints := info.MakeEndpoints(infoService, CompositeRouter)
	complianceEndpoints := compliance.MakeEndpoints(complianceService)
	rulesEndpoints := rules.MakeEndpoints(rulesService)
	return infoEndpoints, complianceEndpoints, rulesEndpoints
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readConfigFile(cfg *Config) {
	f, err := os.Open("config.yml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func readEnvironmentVariables(cfg *Config) {
	port := os.Getenv("SERVER_PORT")
	logLevel := os.Getenv("LOG_LEVEL")
	if port != "" {
		cfg.Server.Port = port
	}
	if logLevel != "" {
		cfg.Log.Level = logLevel
	}
}
