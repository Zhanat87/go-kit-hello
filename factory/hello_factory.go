package factory

import (
	"github.com/Zhanat87/common-libs/instrumenting"
	"github.com/Zhanat87/go-kit-hello/middleware"
	"github.com/Zhanat87/go-kit-hello/service/hello"
	"github.com/go-kit/kit/log"
)

type HelloServiceFactory struct{}

func (s *HelloServiceFactory) CreateHTTPService(logger log.Logger) hello.HTTPService {
	srv := hello.NewHTTPService()
	srv = middleware.NewHelloLoggingMiddleware(log.With(logger, "component", hello.PackageName), srv, hello.PackageName)
	counter, duration, counterError := instrumenting.GetMetricsBySubsystem(hello.PackageName)
	srv = middleware.NewHelloInstrumentingMiddleware(counter, duration, counterError, srv, hello.PackageName)

	return srv
}
