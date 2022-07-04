package main

import (
	"CZERTAINLY-X509-Compliance-Provider/cmd"
	"CZERTAINLY-X509-Compliance-Provider/cmd/attributes"
	"CZERTAINLY-X509-Compliance-Provider/cmd/compliance"
	"CZERTAINLY-X509-Compliance-Provider/cmd/health"
	"CZERTAINLY-X509-Compliance-Provider/cmd/info"
	"CZERTAINLY-X509-Compliance-Provider/cmd/rules"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

// Config is the configuration of the server read from yaml or environment variables
type Config struct {
	Server struct {
		Port     string `yaml:"port", envconfig:"SERVER_PORT"`
		Protocol string `yaml:"protocol", envconfig:"SERVER_PROTOCOL"`
	} `yaml:"server"`
	Log struct {
		Level string `yaml:"level", envconfig:"LOG_LEVEL"`
	} `yaml:"log"`
}

// CompositeRouter is the router for the server
var CompositeRouter = mux.NewRouter()
var config Config
var RULE_FILE_NAME = "./rules.json"
var GROUP_FILE_NAME = "./groups.json"

// init is called before main. It initializes the configuration of the server
func init() {
	readConfigFile(&config)
	readEnvironmentVariables(&config)
}

// main is the entry point of the program
func main() {
	var httpAddr = flag.String(config.Server.Protocol, ":"+config.Server.Port, "http listen address")
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

	// Create a logger
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

	// Create the services
	flag.Parse()
	ctx := context.Background()
	errs := make(chan error)

	infoSrv, complianceSrv, rulesSrv, healthSrv := createService(sugar)
	infoEndpoints, complianceEndpoints, rulesEndpoints, attributeEndPoints, healthEndPoint := createEndPoints(infoSrv, complianceSrv, rulesSrv, healthSrv)

	// Start the services
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := cmd.NewHttpServer(CompositeRouter, ctx, infoEndpoints, complianceEndpoints, rulesEndpoints, attributeEndPoints, healthEndPoint)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()
	fmt.Print(<-errs)
}

// createService creates the services for the main to use it
func createService(logger *zap.SugaredLogger) (info.Service, compliance.Service, rules.Service, health.Service) {
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
		rulesSrv = rules.NewService(logger, RULE_FILE_NAME, GROUP_FILE_NAME)
	}

	var healthSrv health.Service
	{
		healthSrv = health.NewService()
	}

	return infoSrv, complianceSrv, rulesSrv, healthSrv
}

// createEndPoints creates the endpoints for the main to use it
func createEndPoints(infoService info.Service, complianceService compliance.Service, rulesService rules.Service, healthService health.Service) (info.EndPoints, compliance.EndPoints, rules.EndPoints, attributes.EndPoints, health.EndPoints) {
	infoEndpoints := info.MakeEndpoints(infoService, CompositeRouter)
	complianceEndpoints := compliance.MakeEndpoints(complianceService)
	rulesEndpoints := rules.MakeEndpoints(rulesService)
	attributeEndPoint := attributes.MakeEndpoints()
	healthEndPoint := health.MakeEndpoints(healthService)
	return infoEndpoints, complianceEndpoints, rulesEndpoints, attributeEndPoint, healthEndPoint
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

// readConfigFile reads the configuration file from yaml and set them in config
func readConfigFile(cfg *Config) {
	absPath, _ := filepath.Abs("config/config.yml")
	f, err := os.Open(absPath)
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

// readEnvironmentVariables reads the configuration from environment variables and set them in config
func readEnvironmentVariables(cfg *Config) {
	port := os.Getenv("SERVER_PORT")
	logLevel := os.Getenv("LOG_LEVEL")
	protocol := os.Getenv("SERVER_PROTOCOL")
	if port != "" {
		cfg.Server.Port = port
	}
	if logLevel != "" {
		cfg.Log.Level = logLevel
	}
	if protocol != "" {
		cfg.Server.Protocol = protocol
	}
}
