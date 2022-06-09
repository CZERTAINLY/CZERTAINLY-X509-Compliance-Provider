package rules

import (
	"CZERTAINLY-X509-Compliance-Provider/cmd/attributes"
	"context"
	"github.com/gorilla/mux"
	"net/http"
)

type RuleStatus string

const (
	COMPLIANT      RuleStatus = "compliant"
	NON_COMPLIANT  RuleStatus = "nonCompliant"
	NOT_APPLICABLE RuleStatus = "notApplicable"
)

type (
	Request struct {
		Kind            string
		CertificateType []string
	}

	GroupRequest struct {
		UUID string
		Kind string
	}

	Response struct {
		UUID            string                  `json:"uuid"`
		Name            string                  `json:"name"`
		CertificateType string                  `json:"certificateType"`
		Description     string                  `json:"description,omitempty"`
		Attributes      []attributes.Attributes `json:"attributes,omitempty"`
	}

	RuleDefinition struct {
		UUID            string                  `json:"uuid"`
		Name            string                  `json:"name"`
		Description     string                  `json:"description,omitempty"`
		CertificateType string                  `json:"certificateType"`
		Attributes      []attributes.Attributes `json:"attributes,omitempty"`
		GroupUUID       string                  `json:"groupUuid"`
		Kind            string                  `json:"kind"`
	}

	GroupDefinition struct {
		UUID        string `json:"uuid"`
		Description string `json:"description,omitempty"`
		Name        string `json:"name"`
		Kind        string
	}

	GroupResponse struct {
		UUID        string `json:"uuid"`
		Description string `json:"description,omitempty"`
		Name        string `json:"name"`
	}
)

func DecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	kind := vars["kind"]
	return Request{
		kind, nil,
	}, nil
}

func DecodeRuleRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	kind := vars["kind"]
	certificateTypes := r.URL.Query()["certificateType"]
	return Request{
		kind, certificateTypes,
	}, nil
}

func DecodeGroupDetailRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	kind := vars["kind"]
	uuid := vars["uuid"]
	return GroupRequest{
		uuid, kind,
	}, nil
}
