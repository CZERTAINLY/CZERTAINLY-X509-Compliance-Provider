package info

import (
	"github.com/gorilla/mux"
)

type Service interface {
	// GetInfo returns the info for the given request. This information includes the list of end points
	// and the list of attributes.
	GetInfo(router *mux.Router) ([]Response, error)
}
