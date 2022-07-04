package health

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type EndPoints struct {
	GetHealth endpoint.Endpoint
}

func MakeEndpoints(s Service) EndPoints {
	return EndPoints{
		GetHealth: makeHealthEndPoint(s),
	}
}

func makeHealthEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		healthStatus := s.GetHealth()
		return healthStatus, nil
	}
}
