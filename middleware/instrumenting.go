package middleware

import (
	"time"

	"github.com/Zhanat87/common-libs/gokitmiddlewares"
	errorservice "github.com/Zhanat87/go-kit-hello/service/error"
	"github.com/Zhanat87/go-kit-hello/service/hello"
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

func (s *helloInstrumentingMiddleware) Index(req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		s.saver.Save(err, begin, "index")
	}(time.Now())

	return s.next.Index(req)
}

type errorInstrumentingMiddleware struct {
	next  errorservice.HTTPService
	saver gokitmiddlewares.Saver
}

func NewErrorInstrumentingMiddleware(counter metrics.Counter, latency metrics.Histogram,
	counterE metrics.Counter, s errorservice.HTTPService, packageName string) errorservice.HTTPService {
	return &errorInstrumentingMiddleware{
		saver: gokitmiddlewares.NewInstrumenting(counter, latency, counterE, packageName),
		next:  s,
	}
}

func (s *errorInstrumentingMiddleware) Index(req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		s.saver.Save(err, begin, "index")
	}(time.Now())

	return s.next.Index(req)
}
