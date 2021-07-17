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

func NewHelloInstrumentingMiddleware(counter metrics.Counter, latency metrics.Histogram,
	counterE metrics.Counter, s hello.HTTPService, packageName string) hello.HTTPService {
	return &helloInstrumentingMiddleware{
		saver: gokitmiddlewares.NewInstrumenting(counter, latency, counterE, packageName),
		next:  s,
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

func NewErrorInstrumentingMiddleware(counter metrics.Counter, latency metrics.Histogram,
	counterE metrics.Counter, s errorservice.HTTPService, packageName string) errorservice.HTTPService {
	return &errorsInstrumentingMiddleware{
		saver: gokitmiddlewares.NewInstrumenting(counter, latency, counterE, packageName),
		next:  s,
	}
}

func (s *errorsInstrumentingMiddleware) Index(req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		s.saver.Save(err, begin, "index")
	}(time.Now())

	return s.next.Index(req)
}

type pingInstrumentingMiddleware struct {
	next  ping.HTTPService
	saver gokitmiddlewares.Saver
}

func NewPingInstrumentingMiddleware(counter metrics.Counter, latency metrics.Histogram,
	counterE metrics.Counter, s ping.HTTPService, packageName string) ping.HTTPService {
	return &pingInstrumentingMiddleware{
		saver: gokitmiddlewares.NewInstrumenting(counter, latency, counterE, packageName),
		next:  s,
	}
}

func (s *pingInstrumentingMiddleware) Grpc(req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		s.saver.Save(err, begin, "grpc")
	}(time.Now())

	return s.next.Grpc(req)
}

func (s *pingInstrumentingMiddleware) HTTP(req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		s.saver.Save(err, begin, "http")
	}(time.Now())

	return s.next.HTTP(req)
}
