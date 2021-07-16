package factory

import (
	"github.com/Zhanat87/common-libs/instrumenting"
	"github.com/Zhanat87/go-kit-hello/middleware"
	"github.com/Zhanat87/go-kit-hello/service/ping"
	"github.com/go-kit/kit/log"
)

type PingServiceFactory struct{}

func (s *PingServiceFactory) CreateHTTPService(logger log.Logger) ping.HTTPService {
	srv := ping.NewHTTPService()
	srv = middleware.NewPingLoggingMiddleware(log.With(logger, "component", ping.PackageName), srv, ping.PackageName)
	counter, duration, counterPing := instrumenting.GetMetricsBySubsystem(ping.PackageName)
	srv = middleware.NewPingInstrumentingMiddleware(counter, duration, counterPing, srv, ping.PackageName)

	return srv
}
