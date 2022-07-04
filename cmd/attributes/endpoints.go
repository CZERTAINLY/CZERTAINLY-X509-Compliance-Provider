package attributes

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type EndPoints struct {
	GetAttributes      endpoint.Endpoint
	ValidateAttributes endpoint.Endpoint
}

func MakeEndpoints() EndPoints {
	return EndPoints{
		GetAttributes:      makeAttributeEndpoint(),
		ValidateAttributes: makeValidateAttributeEndpoint(),
	}
}

func makeAttributeEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return []string{}, nil
	}
}

func makeValidateAttributeEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return nil, nil
	}
}
