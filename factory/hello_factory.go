package factory

import (
	"github.com/Zhanat87/common-libs/instrumenting"
	"github.com/Zhanat87/go-kit-hello/middleware"
	"github.com/Zhanat87/go-kit-hello/service/hello"
	"github.com/go-kit/kit/log"
	"github.com/openzipkin/zipkin-go"
)

type HelloServiceFactory struct{}

func (s *HelloServiceFactory) CreateHTTPService(packageName string, logger log.Logger, zipkinTracer *zipkin.Tracer) hello.HTTPService {
	srv := hello.NewHTTPService()
	srv = middleware.NewHelloLoggingMiddleware(srv, packageName, log.With(logger, "component", packageName))
	counter, duration, counterError := instrumenting.GetMetricsBySubsystem(packageName)
	srv = middleware.NewHelloInstrumentingMiddleware(srv, packageName, counter, duration, counterError)
	srv = middleware.NewHelloZipkinTracingMiddleware(srv, packageName, zipkinTracer)

	return srv
}
