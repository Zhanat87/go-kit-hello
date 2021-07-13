package service

import (
	"github.com/Zhanat87/go-kit-hello/contracts"
	"github.com/Zhanat87/go-kit-hello/domain"
)

type helloService struct{}

func NewHelloService() contracts.HelloService {
	return &helloService{}
}

func (s *helloService) SayHi(name string) string {
	model := domain.Model{
		Name: name,
	}

	return model.SayHi()
}
