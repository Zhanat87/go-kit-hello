package service

import (
	"errors"

	"github.com/Zhanat87/go-kit-hello/contracts"
	"github.com/Zhanat87/go-kit-hello/transport"
)

type helloHTTPService struct {
	helloService contracts.HelloService
}

func NewHTTPService() contracts.HTTPService {
	return &helloHTTPService{helloService: NewHelloService()}
}

func (s *helloHTTPService) Index(req interface{}) (interface{}, error) {
	r := req.(*transport.HelloRequest)

	return &transport.HelloResponse{Data: s.helloService.SayHi(r.Name)}, nil
}

func (s *helloHTTPService) Error(req interface{}) (interface{}, error) {
	return &transport.HelloResponse{Data: "error response"}, errors.New("error from hello")
}
