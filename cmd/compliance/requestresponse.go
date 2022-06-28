package compliance

import (
	"CZERTAINLY-X509-Compliance-Provider/cmd/attributes"
	"CZERTAINLY-X509-Compliance-Provider/cmd/rules"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Status string

const (
	OK  Status = "ok"
	NOK Status = "nok"
	NA  Status = "na"
)

type (
	RequestRules struct {
		UUID       string                         `json:"uuid"`
		Attributes []attributes.RequestAttributes `json:"attributes,omitempty"`
	}

	Request struct {
		Certificate string         `json:"certificate"`
		Rules       []RequestRules `json:"rules"`
	}

	RequestAndKind struct {
		Request Request
		Kind    string
	}

	ResponseRules struct {
		UUID   string           `json:"uuid"`
		Name   string           `json:"name"`
		Status rules.RuleStatus `json:"status"`
	}

	Response struct {
		Status Status          `json:"status"`
		Rules  []ResponseRules `json:"rules"`
	}
)

func DecodeComplianceRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req Request
	vars := mux.Vars(r)
	kind := vars["kind"]
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return RequestAndKind{
		req, kind,
	}, nil
}
