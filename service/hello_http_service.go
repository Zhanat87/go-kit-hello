package service

import (
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
