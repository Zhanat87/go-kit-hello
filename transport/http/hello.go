package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Zhanat87/go-kit-hello/encoders"
	"github.com/Zhanat87/go-kit-hello/middleware"
	"github.com/Zhanat87/go-kit-hello/transport"
	kitlog "github.com/go-kit/kit/log"
	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandler(srvEndpoints middleware.Endpoints, logger kitlog.Logger,
	baseURL string, decodeIndexRequestFunc kithttp.DecodeRequestFunc) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encoders.EncodeErrorJSON),
	}
	index := kithttp.NewServer(
		srvEndpoints.IndexEndpoint,
		decodeIndexRequestFunc,
		encoders.EncodeResponseJSON,
		opts...,
	)
	r := mux.NewRouter()
	r.Handle(baseURL, index).Methods(http.MethodPost)
	r.Handle(baseURL, index).Methods(http.MethodGet)

	return r
}

func DecodeIndexRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body transport.HelloRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return body, nil
}
