package cmd

import (
	"CZERTAINLY-X509-Compliance-Provider/cmd/attributes"
	"CZERTAINLY-X509-Compliance-Provider/cmd/compliance"
	"CZERTAINLY-X509-Compliance-Provider/cmd/health"
	"CZERTAINLY-X509-Compliance-Provider/cmd/info"
	"CZERTAINLY-X509-Compliance-Provider/cmd/rules"
	"CZERTAINLY-X509-Compliance-Provider/cmd/utils"
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHttpServer(router *mux.Router, ctx context.Context, infoEndpoint info.EndPoints,
	complianceEndpoints compliance.EndPoints, rulesEndPoints rules.EndPoints,
	attributeEndPoints attributes.EndPoints, healthEndPoint health.EndPoints) http.Handler {
	router.Use(commonMiddleware)

	router.Methods("GET").Path("/v1").Handler(httptransport.NewServer(
		infoEndpoint.GetInfo, utils.DecodeRequest, utils.EncodeResponse,
	)).Name("listSupportedFunctions")

	router.Methods("GET").Path("/v1/complianceProvider/{kind}/attributes").Handler(httptransport.NewServer(
		attributeEndPoints.GetAttributes, utils.DecodeRequest, utils.EncodeResponse,
	)).Name("listAttributes")

	router.Methods("POST").Path("/v1/complianceProvider/{kind}/attributes").Handler(httptransport.NewServer(
		attributeEndPoints.ValidateAttributes, utils.DecodeRequest, utils.EncodeResponse,
	)).Name("validateAttributes")

	router.Methods("GET").Path("/v1/health").Handler(httptransport.NewServer(
		healthEndPoint.GetHealth, utils.DecodeRequest, utils.EncodeResponse,
	)).Name("getHealth")

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
