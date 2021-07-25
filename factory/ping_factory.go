package factory

import (
	"github.com/Zhanat87/common-libs/instrumenting"
	"github.com/Zhanat87/go-kit-hello/middleware"
	"github.com/Zhanat87/go-kit-hello/service/ping"
	"github.com/go-kit/kit/log"
	"github.com/openzipkin/zipkin-go"
)

type PingServiceFactory struct{}

func (s *PingServiceFactory) CreateHTTPService(packageName string, logger log.Logger, zipkinTracer *zipkin.Tracer) ping.HTTPService {
	srv := ping.NewHTTPService(zipkinTracer)
	srv = middleware.NewPingLoggingMiddleware(srv, packageName, log.With(logger, "component", packageName))
	counter, duration, counterError := instrumenting.GetMetricsBySubsystem(packageName)
	srv = middleware.NewPingInstrumentingMiddleware(srv, packageName, counter, duration, counterError)
	srv = middleware.NewPingZipkinTracingMiddleware(srv, packageName, zipkinTracer)

	return srv
}
