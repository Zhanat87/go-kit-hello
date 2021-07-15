package middleware

import (
	"time"

	"github.com/Zhanat87/go-kit-hello/contracts"
	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	requestError   metrics.Counter
	next           contracts.HTTPService
	packageName    string
}

func NewInstrumentingMiddleware(counter metrics.Counter, latency metrics.Histogram,
	counterE metrics.Counter, s contracts.HTTPService, packageName string) contracts.HTTPService {
	return &instrumentingMiddleware{
		requestCount:   counter,
		requestLatency: latency,
		requestError:   counterE,
		next:           s,
		packageName:    packageName,
	}
}

func (s *instrumentingMiddleware) Index(req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		println("a")
		labels := []string{"method", s.packageName + "_index"}
		s.requestCount.With(labels...).Add(1)
		if err != nil {
			println("b")
			s.requestError.With(labels...).Add(1)
		}
		println("c")
		s.requestLatency.With(labels...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.next.Index(req)
}

func (s *instrumentingMiddleware) Error(req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		labels := []string{"method", s.packageName + "_error"}
		s.requestCount.With(labels...).Add(1)
		if err != nil {
			println("hello error")
			s.requestError.With(labels...).Add(1)
		}
		s.requestLatency.With(labels...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.next.Error(req)
}

func (s *instrumentingMiddleware) Grpc(req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		labels := []string{"method", s.packageName + "_grpc"}
		s.requestCount.With(labels...).Add(1)
		if err != nil {
			println("hello grpc")
			s.requestError.With(labels...).Add(1)
		}
		s.requestLatency.With(labels...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.next.Grpc(req)
}
