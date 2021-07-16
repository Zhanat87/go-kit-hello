package factory

import (
	"github.com/Zhanat87/common-libs/instrumenting"
	"github.com/Zhanat87/go-kit-hello/middleware"
	"github.com/Zhanat87/go-kit-hello/service/error"
	"github.com/go-kit/kit/log"
)

type ErrorServiceFactory struct{}

func (s *ErrorServiceFactory) CreateHTTPService(logger log.Logger) error.HTTPService {
	srv := error.NewHTTPService()
	srv = middleware.NewErrorLoggingMiddleware(log.With(logger, "component", error.PackageName), srv, error.PackageName)
	counter, duration, counterError := instrumenting.GetMetricsBySubsystem(error.PackageName)
	srv = middleware.NewErrorInstrumentingMiddleware(counter, duration, counterError, srv, error.PackageName)

	return srv
}
