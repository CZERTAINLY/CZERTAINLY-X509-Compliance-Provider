package info

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

type service struct {
	logger log.Logger
}

var endpoints []EndPointsInfo

func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
	}
}

func (s service) GetInfo(ctx context.Context, router *mux.Router) (Response, error) {
	//logger := log.With(s.logger, "method", "GetInfo")
	endpoints = endpoints[:0]
	router.Walk(gorillaWalkFn)
	infoResponse := Response{
		FunctionGroupCode: "complianceProvider",
		Kinds:             []string{"x509"},
		EndPoints:         endpoints,
	}
	return infoResponse, nil
}

func gorillaWalkFn(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	path, _ := route.GetPathTemplate()
	method, _ := route.GetMethods()
	name := route.GetName()
	endpoints = append(endpoints,
		EndPointsInfo{
			Name:    name,
			Context: path,
			Method:  method[0],
		})
	return nil
}
