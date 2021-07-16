package ping

import "github.com/Zhanat87/go-kit-hello/transport"

type HTTPService interface {
	Grpc(req interface{}) (response interface{}, err error)
	HTTP(req interface{}) (response interface{}, err error)
}

type httpService struct{}

func NewHTTPService() HTTPService {
	return &httpService{}
}

func (s *httpService) Grpc(req interface{}) (interface{}, error) {
	return &transport.PingRequest{Ping: "grpc pong response"}, nil
}

func (s *httpService) HTTP(req interface{}) (interface{}, error) {
	return &transport.PingRequest{Ping: "http pong response"}, nil
}
