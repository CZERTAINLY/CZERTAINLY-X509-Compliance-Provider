package info

import (
	"github.com/gorilla/mux"
)

type Service interface {
	GetInfo(router *mux.Router) ([]Response, error)
}
