package rules

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type EndPoints struct {
	GetRules       endpoint.Endpoint
	GetGroups      endpoint.Endpoint
	GetGroupDetail endpoint.Endpoint
}

func MakeEndpoints(s Service) EndPoints {
	return EndPoints{
		GetRules:       makeGetRulesEndpoint(s),
		GetGroups:      makeGetGroupsEndpoint(s),
		GetGroupDetail: makeGetGroupDetailsEndpoint(s),
	}
}

func makeGetRulesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Request)
		response, err := s.GetRules(ctx, req.Kind, req.CertificateType)
		return response, err
	}
}

func makeGetGroupsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Request)
		response, err := s.GetGroups(ctx, req.Kind)
		return response, err
	}
}

func makeGetGroupDetailsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GroupRequest)
		response, err := s.GetGroupDetails(ctx, req.UUID, req.Kind)
		return response, err
	}
}
