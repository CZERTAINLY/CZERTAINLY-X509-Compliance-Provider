package info

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
)

type EndPoints struct {
	GetInfo endpoint.Endpoint
}

func MakeEndpoints(s Service, router *mux.Router) EndPoints {
	return EndPoints{
		GetInfo: makeGetInfoEndpoint(s, router),
	}
}

func makeGetInfoEndpoint(s Service, router *mux.Router) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		infoResponse, err := s.GetInfo(router)
		return infoResponse, err
	}
}
