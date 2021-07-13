package factory

import (
	"github.com/Zhanat87/go-kit-hello/contracts"
	"github.com/Zhanat87/go-kit-hello/instrumenting"
	"github.com/Zhanat87/go-kit-hello/middleware"
	"github.com/Zhanat87/go-kit-hello/service"
	"github.com/Zhanat87/go-kit-hello/utils"
	"github.com/go-kit/kit/log"
)

type ServiceFactory struct{}

func (s *ServiceFactory) CreateHTTPService(logger log.Logger) contracts.HTTPService {
	srv := service.NewHTTPService()
	srv = middleware.NewLoggingMiddleware(log.With(logger, "component", utils.PackageName), srv, utils.PackageName)
	counter, duration, counterError := instrumenting.GetMetricsBySubsystem(utils.PackageName)
	srv = middleware.NewInstrumentingMiddleware(counter, duration, counterError, srv, utils.PackageName)

	return srv
}