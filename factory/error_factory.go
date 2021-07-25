package factory

import (
	"github.com/Zhanat87/common-libs/instrumenting"
	"github.com/Zhanat87/go-kit-hello/middleware"
	"github.com/Zhanat87/go-kit-hello/service/error"
	"github.com/go-kit/kit/log"
	"github.com/openzipkin/zipkin-go"
)

type ErrorServiceFactory struct{}

func (s *ErrorServiceFactory) CreateHTTPService(packageName string, logger log.Logger, zipkinTracer *zipkin.Tracer) error.HTTPService {
	srv := error.NewHTTPService()
	srv = middleware.NewErrorLoggingMiddleware(srv, packageName, log.With(logger, "component", packageName))
	counter, duration, counterError := instrumenting.GetMetricsBySubsystem(packageName)
	srv = middleware.NewErrorInstrumentingMiddleware(srv, packageName, counter, duration, counterError)
	srv = middleware.NewErrorZipkinTracingMiddleware(srv, packageName, zipkinTracer)

	return srv
}
