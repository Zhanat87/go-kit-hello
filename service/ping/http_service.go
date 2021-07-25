package ping

import (
	"context"
	"fmt"

	"github.com/Zhanat87/go-kit-hello/transport"
)

type HTTPService interface {
	Grpc(ctx context.Context, req interface{}) (response interface{}, err error)
	HTTP(ctx context.Context, req interface{}) (response interface{}, err error)
}

type httpService struct {
	service Service
}

func NewHTTPService() HTTPService {
	return &httpService{service: NewService()}
}

func (s *httpService) Grpc(ctx context.Context, req interface{}) (interface{}, error) {
	return &transport.PingRequest{Ping: "grpc pong response"}, nil
}

func (s *httpService) HTTP(ctx context.Context, req interface{}) (interface{}, error) {
	pingRequest, ok := req.(*transport.PingRequest)
	if !ok {
		return nil, fmt.Errorf("error convert transport.PingRequest: %#v", req)
	}
	pong, err := s.service.Ping(ctx, pingRequest.Ping)

	return &transport.PingResponse{Pong: pong}, err
}
