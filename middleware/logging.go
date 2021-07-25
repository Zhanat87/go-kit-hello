package middleware

import (
	"context"
	"time"

	"github.com/Zhanat87/common-libs/gokitmiddlewares"
	errorservice "github.com/Zhanat87/go-kit-hello/service/error"
	"github.com/Zhanat87/go-kit-hello/service/hello"
	"github.com/Zhanat87/go-kit-hello/service/ping"
	"github.com/go-kit/kit/log"
)

type helloLoggingMiddleware struct {
	next  hello.HTTPService
	saver gokitmiddlewares.Saver
}

func NewHelloLoggingMiddleware(s hello.HTTPService, packageName string, logger log.Logger) hello.HTTPService {
	return &helloLoggingMiddleware{
		next:  s,
		saver: gokitmiddlewares.NewLogging(logger, packageName),
	}
}

func (s *helloLoggingMiddleware) Index(ctx context.Context, req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		s.saver.Save(err, begin, "index")
	}(time.Now())

	return s.next.Index(ctx, req)
}

type errorLoggingMiddleware struct {
	next  errorservice.HTTPService
	saver gokitmiddlewares.Saver
}

func NewErrorLoggingMiddleware(s errorservice.HTTPService, packageName string, logger log.Logger) errorservice.HTTPService {
	return &errorLoggingMiddleware{
		next:  s,
		saver: gokitmiddlewares.NewLogging(logger, packageName),
	}
}

func (s *errorLoggingMiddleware) Index(ctx context.Context, req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		s.saver.Save(err, begin, "index")
	}(time.Now())

	return s.next.Index(ctx, req)
}

type pingLoggingMiddleware struct {
	next  ping.HTTPService
	saver gokitmiddlewares.Saver
}

func NewPingLoggingMiddleware(s ping.HTTPService, packageName string, logger log.Logger) ping.HTTPService {
	return &pingLoggingMiddleware{
		next:  s,
		saver: gokitmiddlewares.NewLogging(logger, packageName),
	}
}

func (s *pingLoggingMiddleware) HTTP(ctx context.Context, req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		s.saver.Save(err, begin, "http")
	}(time.Now())

	return s.next.HTTP(ctx, req)
}

func (s *pingLoggingMiddleware) Grpc(ctx context.Context, req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		s.saver.Save(err, begin, "grpc")
	}(time.Now())

	return s.next.Grpc(ctx, req)
}
