package info

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type service struct {
	logger *zap.SugaredLogger
}

var endpoints []EndPointsInfo

func NewService(logger *zap.SugaredLogger) Service {
	return &service{
		logger: logger,
	}
}

func (s service) GetInfo(router *mux.Router) ([]Response, error) {
	s.logger.Info("Entering GetInfo method")
	endpoints = endpoints[:0]
	err := router.Walk(s.gorillaWalkFn)
	if err != nil {
		return nil, err
	}
	var infoResponse []Response
	infoResponse = append(infoResponse, Response{
		FunctionGroupCode: "complianceProvider",
		Kinds:             []string{"x509"},
		EndPoints:         endpoints,
	})
	s.logger.Info("List of endpoints for the connector is ", infoResponse)
	return infoResponse, nil
}

func (s service) gorillaWalkFn(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	s.logger.Info("Entering gorillaWalkFn to calculate the list of end points")
	path, _ := route.GetPathTemplate()
	s.logger.Debug("Path: ", path)
	method, _ := route.GetMethods()
	s.logger.Debug("Method: ", method)
	name := route.GetName()
	s.logger.Debug("Name: ", name)
	endpoint := EndPointsInfo{
		Name:    name,
		Context: path,
		Method:  method[0],
	}
	s.logger.Debug("Endpoint created: ", endpoint)
	endpoints = append(endpoints,
		endpoint)
	return nil
}
