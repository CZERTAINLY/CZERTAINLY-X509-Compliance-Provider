package info

import (
	"context"
	"github.com/gorilla/mux"
)

type Service interface {
	GetInfo(ctx context.Context, router *mux.Router) ([]Response, error)
}
