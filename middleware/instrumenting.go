package middleware

import (
	"context"
	"time"

	"github.com/Zhanat87/common-libs/gokitmiddlewares"
	errorservice "github.com/Zhanat87/go-kit-hello/service/error"
	"github.com/Zhanat87/go-kit-hello/service/hello"
	"github.com/Zhanat87/go-kit-hello/service/ping"
	"github.com/go-kit/kit/metrics"
)

type helloInstrumentingMiddleware struct {
	next  hello.HTTPService
	saver gokitmiddlewares.Saver
}

func NewHelloInstrumentingMiddleware(s hello.HTTPService, packageName string,
	counter metrics.Counter, latency metrics.Histogram, counterE metrics.Counter) hello.HTTPService {
	return &helloInstrumentingMiddleware{
		next:  s,
		saver: gokitmiddlewares.NewInstrumenting(counter, latency, counterE, packageName),
	}
}

func (s *helloInstrumentingMiddleware) Index(ctx context.Context, req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		s.saver.Save(err, begin, "index")
	}(time.Now())

	return s.next.Index(ctx, req)
}

type errorsInstrumentingMiddleware struct {
	next  errorservice.HTTPService
	saver gokitmiddlewares.Saver
}

func NewErrorInstrumentingMiddleware(s errorservice.HTTPService, packageName string,
	counter metrics.Counter, latency metrics.Histogram, counterE metrics.Counter) errorservice.HTTPService {
	return &errorsInstrumentingMiddleware{
		next:  s,
		saver: gokitmiddlewares.NewInstrumenting(counter, latency, counterE, packageName),
	}
}

func (s *errorsInstrumentingMiddleware) Index(ctx context.Context, req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		s.saver.Save(err, begin, "index")
	}(time.Now())

	return s.next.Index(ctx, req)
}

type pingInstrumentingMiddleware struct {
	next  ping.HTTPService
	saver gokitmiddlewares.Saver
}

func NewPingInstrumentingMiddleware(s ping.HTTPService, packageName string,
	counter metrics.Counter, latency metrics.Histogram, counterE metrics.Counter) ping.HTTPService {
	return &pingInstrumentingMiddleware{
		next:  s,
		saver: gokitmiddlewares.NewInstrumenting(counter, latency, counterE, packageName),
	}
}

func (s *pingInstrumentingMiddleware) Grpc(ctx context.Context, req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		s.saver.Save(err, begin, "grpc")
	}(time.Now())

	return s.next.Grpc(ctx, req)
}

func (s *pingInstrumentingMiddleware) HTTP(ctx context.Context, req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		s.saver.Save(err, begin, "http")
	}(time.Now())

	return s.next.HTTP(ctx, req)
}
