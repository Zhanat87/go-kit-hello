package ping

import (
	"context"
	"fmt"

	"github.com/openzipkin/zipkin-go"

	"github.com/Zhanat87/go-kit-hello/transport"
)

type HTTPService interface {
	Grpc(ctx context.Context, req interface{}) (response interface{}, err error)
	HTTP(ctx context.Context, req interface{}) (response interface{}, err error)
}

type httpService struct {
	service      Service
	zipkinTracer *zipkin.Tracer
}

func NewHTTPService(zipkinTracer *zipkin.Tracer) HTTPService {
	return &httpService{service: NewService(zipkinTracer), zipkinTracer: zipkinTracer}
}

func (s *httpService) Grpc(ctx context.Context, req interface{}) (interface{}, error) {
	return &transport.PingRequest{Ping: "grpc pong response"}, nil
}

func (s *httpService) HTTP(ctx context.Context, req interface{}) (interface{}, error) {
	span, ctx := s.zipkinTracer.StartSpanFromContext(ctx, "ping http service span")
	defer span.Finish()
	pingRequest, ok := req.(*transport.PingRequest)
	if !ok {
		return nil, fmt.Errorf("error convert transport.PingRequest: %#v", req)
	}
	pong, err := s.service.Ping(ctx, pingRequest.Ping)
	if err != nil {
		return nil, err
	}

	return &transport.PingResponse{Pong: pong}, nil
}
