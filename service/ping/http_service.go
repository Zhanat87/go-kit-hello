package ping

import (
	"context"

	"github.com/Zhanat87/go-kit-hello/transport"
)

type HTTPService interface {
	Grpc(ctx context.Context, req interface{}) (response interface{}, err error)
	HTTP(ctx context.Context, req interface{}) (response interface{}, err error)
}

type httpService struct{}

func NewHTTPService() HTTPService {
	return &httpService{}
}

func (s *httpService) Grpc(ctx context.Context, req interface{}) (interface{}, error) {
	return &transport.PingRequest{Ping: "grpc pong response"}, nil
}

func (s *httpService) HTTP(ctx context.Context, req interface{}) (interface{}, error) {
	return &transport.PingRequest{Ping: "http pong response"}, nil
}
