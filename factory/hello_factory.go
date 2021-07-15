package factory

import (
	"github.com/Zhanat87/go-kit-hello/contracts"
	"github.com/Zhanat87/go-kit-hello/instrumenting"
	"github.com/Zhanat87/go-kit-hello/middleware"
	"github.com/Zhanat87/go-kit-hello/service"
	"github.com/Zhanat87/go-kit-hello/utils"
	"github.com/go-kit/kit/log"
	"github.com/openzipkin/zipkin-go"
)

type ServiceFactory struct{}

func (s *ServiceFactory) CreateHTTPService(logger log.Logger, tracer *zipkin.Tracer) contracts.HTTPService {
	srv := service.NewHTTPService(tracer)
	srv = middleware.NewLoggingMiddleware(log.With(logger, "component", utils.PackageName), srv, utils.PackageName)
	counter, duration, counterError := instrumenting.GetMetricsBySubsystem(utils.PackageName)
	srv = middleware.NewInstrumentingMiddleware(counter, duration, counterError, srv, utils.PackageName)

	return srv
}
