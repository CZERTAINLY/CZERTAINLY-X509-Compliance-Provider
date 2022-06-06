package utils

import (
	"context"
	"encoding/json"
	"net/http"
)

func EncodeResponse(ctx context.Context, writer http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(writer).Encode(response)
}

func DecodeRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	return request, nil
}
