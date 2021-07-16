package middleware

import (
	"time"

	"github.com/Zhanat87/common-libs/gokitmiddlewares"
	"github.com/Zhanat87/go-kit-hello/service/hello"
	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	next  hello.HTTPService
	saver gokitmiddlewares.Saver
}

func NewInstrumentingMiddleware(counter metrics.Counter, latency metrics.Histogram,
	counterE metrics.Counter, s hello.HTTPService, packageName string) hello.HTTPService {
	return &instrumentingMiddleware{
		saver: gokitmiddlewares.NewInstrumenting(counter, latency, counterE, packageName),
		next:  s,
	}
}

func (s *instrumentingMiddleware) Index(req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		s.saver.Save(err, begin, "index")
	}(time.Now())

	return s.next.Index(req)
}
