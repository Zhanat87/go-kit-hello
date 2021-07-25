package http

import (
	"context"
	"encoding/json"
	"net/http"
	"github.com/Zhanat87/common-libs/tracers"

	"github.com/Zhanat87/common-libs/encoders"
	"github.com/Zhanat87/common-libs/gokithttp"
	"github.com/Zhanat87/go-kit-hello/middleware"
	"github.com/Zhanat87/go-kit-hello/transport"
	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHelloHandler(srvEndpoints middleware.HelloEndpoints,
	logger kitlog.Logger, baseURL string) http.Handler {
	opts := gokithttp.GetServerOptionsWithZipkinTracer(logger, tracers.ZipkinTracer)
	index := kithttp.NewServer(
		srvEndpoints.IndexEndpoint,
		DecodeHelloIndexRequest,
		encoders.EncodeResponseJSON,
		opts...,
	)
	r := mux.NewRouter()
	r.Handle(baseURL+"index", index).Methods(http.MethodPost)

	return r
}

func DecodeHelloIndexRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body transport.HelloRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return body, nil
}
