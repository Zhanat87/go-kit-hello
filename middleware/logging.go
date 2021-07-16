package middleware

import (
	"time"

	"github.com/Zhanat87/common-libs/gokitmiddlewares"
	"github.com/Zhanat87/go-kit-hello/service/hello"
	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	saver gokitmiddlewares.Saver
	next  hello.HTTPService
}

func NewLoggingMiddleware(logger log.Logger, s hello.HTTPService, packageName string) hello.HTTPService {
	return &loggingMiddleware{
		saver: gokitmiddlewares.NewLogging(logger, packageName),
		next:  s,
	}
}

func (s *loggingMiddleware) Index(req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		s.saver.Save(err, begin, "index")
	}(time.Now())

	return s.next.Index(req)
}
