package middleware

import (
	"time"

	"github.com/Zhanat87/common-libs/gokitmiddlewares"
	errorservice "github.com/Zhanat87/go-kit-hello/service/error"
	"github.com/Zhanat87/go-kit-hello/service/hello"
	"github.com/Zhanat87/go-kit-hello/service/ping"
	"github.com/go-kit/kit/log"
)

type helloLoggingMiddleware struct {
	saver gokitmiddlewares.Saver
	next  hello.HTTPService
}

func NewHelloLoggingMiddleware(logger log.Logger,
	s hello.HTTPService, packageName string) hello.HTTPService {
	return &helloLoggingMiddleware{
		saver: gokitmiddlewares.NewLogging(logger, packageName),
		next:  s,
	}
}

func (s *helloLoggingMiddleware) Index(req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		s.saver.Save(err, begin, "index")
	}(time.Now())

	return s.next.Index(req)
}

type errorLoggingMiddleware struct {
	saver gokitmiddlewares.Saver
	next  errorservice.HTTPService
}

func NewErrorLoggingMiddleware(logger log.Logger,
	s errorservice.HTTPService, packageName string) errorservice.HTTPService {
	return &errorLoggingMiddleware{
		saver: gokitmiddlewares.NewLogging(logger, packageName),
		next:  s,
	}
}

func (s *errorLoggingMiddleware) Index(req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		s.saver.Save(err, begin, "index")
	}(time.Now())

	return s.next.Index(req)
}

type pingLoggingMiddleware struct {
	saver gokitmiddlewares.Saver
	next  ping.HTTPService
}

func NewPingLoggingMiddleware(logger log.Logger,
	s ping.HTTPService, packageName string) ping.HTTPService {
	return &pingLoggingMiddleware{
		saver: gokitmiddlewares.NewLogging(logger, packageName),
		next:  s,
	}
}

func (s *pingLoggingMiddleware) Grpc(req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		s.saver.Save(err, begin, "grpc")
	}(time.Now())

	return s.next.Grpc(req)
}

func (s *pingLoggingMiddleware) HTTP(req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		s.saver.Save(err, begin, "http")
	}(time.Now())

	return s.next.HTTP(req)
}
