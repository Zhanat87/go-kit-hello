package factory

import (
	"github.com/Zhanat87/common-libs/instrumenting"
	"github.com/Zhanat87/go-kit-hello/middleware"
	"github.com/Zhanat87/go-kit-hello/service/hello"
	"github.com/go-kit/kit/log"
	"github.com/openzipkin/zipkin-go"
)

type ServiceFactory struct{}

func (s *ServiceFactory) CreateHTTPService(logger log.Logger, tracer *zipkin.Tracer) hello.HTTPService {
	srv := hello.NewHTTPService(tracer)
	srv = middleware.NewLoggingMiddleware(log.With(logger, "component", hello.PackageName), srv, hello.PackageName)
	counter, duration, counterError := instrumenting.GetMetricsBySubsystem(hello.PackageName)
	srv = middleware.NewInstrumentingMiddleware(counter, duration, counterError, srv, hello.PackageName)

	return srv
}
