package hello

import (
	"context"

	"github.com/Zhanat87/common-libs/tracers"
	"github.com/Zhanat87/go-kit-hello/transport"
)

type HTTPService interface {
	Index(ctx context.Context, req interface{}) (response interface{}, err error)
}

type httpService struct {
	service Service
}

func NewHTTPService() HTTPService {
	return &httpService{service: NewService()}
}

func (s *httpService) Index(ctx context.Context, req interface{}) (interface{}, error) {
	// utils.PrintContextInternals("httpService", ctx, false)
	tracers.ZipkinTracer.StartSpanFromContext(
		ctx,
		"httpService",
	)
	r := req.(*transport.HelloRequest)

	return &transport.HelloResponse{Data: s.service.SayHi(ctx, r.Name)}, nil
}
