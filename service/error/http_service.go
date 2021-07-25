package error

import (
	"context"
	"errors"

	"github.com/Zhanat87/go-kit-hello/transport"
)

type HTTPService interface {
	Index(ctx context.Context, req interface{}) (response interface{}, err error)
}

type httpService struct{}

func NewHTTPService() HTTPService {
	return &httpService{}
}

func (s *httpService) Index(ctx context.Context, req interface{}) (interface{}, error) {
	return &transport.ErrorResponse{Error: "error response"}, errors.New("error text")
}
