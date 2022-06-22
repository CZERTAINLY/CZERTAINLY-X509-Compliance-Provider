package compliance

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type EndPoints struct {
	ComplianceCheck endpoint.Endpoint
}

func MakeEndpoints(s Service) EndPoints {
	return EndPoints{
		ComplianceCheck: makeComplianceCheckEndpoint(s),
	}
}

func makeComplianceCheckEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RequestAndKind)
		response, err := s.ComplianceCheck(req.Kind, req.Request)
		return response, err
	}
}
