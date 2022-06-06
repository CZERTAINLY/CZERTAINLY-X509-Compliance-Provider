package main

import (
	"CZERTAINLY-X509-Compliance-Provider/cmd"
	"CZERTAINLY-X509-Compliance-Provider/cmd/compliance"
	"CZERTAINLY-X509-Compliance-Provider/cmd/info"
	"CZERTAINLY-X509-Compliance-Provider/cmd/rules"
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var CompositeRouter = mux.NewRouter()

func main() {
	var httpAddr = flag.String("http", ":10180", "http listen address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	flag.Parse()
	ctx := context.Background()
	errs := make(chan error)

	infoSrv, complianceSrv, rulesSrv := createService(logger)
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

func createService(logger log.Logger) (info.Service, compliance.Service, rules.Service) {
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
