package cmd

import (
	"CZERTAINLY-X509-Compliance-Provider/cmd/compliance"
	"CZERTAINLY-X509-Compliance-Provider/cmd/info"
	"CZERTAINLY-X509-Compliance-Provider/cmd/rules"
	"CZERTAINLY-X509-Compliance-Provider/cmd/utils"
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHttpServer(router *mux.Router, ctx context.Context, infoEndpoint info.EndPoints, complianceEndpoints compliance.EndPoints, rulesEndPoints rules.EndPoints) http.Handler {
	router.Use(commonMiddleware)

	router.Methods("GET").Path("/v1").Handler(httptransport.NewServer(
		infoEndpoint.GetInfo, utils.DecodeRequest, utils.EncodeResponse,
	)).Name("listSupportedFunctions")

	router.Methods("POST").Path("/v1/complianceProvider/{kind}/compliance").Handler(httptransport.NewServer(
		complianceEndpoints.ComplianceCheck, compliance.DecodeComplianceRequest, utils.EncodeResponse,
	)).Name("checkCompliance")

	router.Methods("GET").Path("/v1/complianceProvider/{kind}/rules").Handler(httptransport.NewServer(
		rulesEndPoints.GetRules, rules.DecodeRuleRequest, utils.EncodeResponse,
	)).Name("listRules")

	router.Methods("GET").Path("/v1/complianceProvider/{kind}/groups").Handler(httptransport.NewServer(
		rulesEndPoints.GetGroups, rules.DecodeRequest, utils.EncodeResponse,
	)).Name("listGroups")

	router.Methods("GET").Path("/v1/complianceProvider/{kind}/groups/{uuid}").Handler(httptransport.NewServer(
		rulesEndPoints.GetGroupDetail, rules.DecodeGroupDetailRequest, utils.EncodeResponse,
	)).Name("listGroupDetails")

	return router
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(writer, request)
	})
}
