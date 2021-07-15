package middleware

import (
	"time"

	"github.com/Zhanat87/go-kit-hello/contracts"
	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger      log.Logger
	next        contracts.HTTPService
	packageName string
}

func NewLoggingMiddleware(logger log.Logger, s contracts.HTTPService, packageName string) contracts.HTTPService {
	return &loggingMiddleware{logger, s, packageName}
}

func (s *loggingMiddleware) Index(req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		println("e")
		if err != nil {
			println("f")
			_ = s.logger.Log(
				"method", s.packageName+"_index",
				"took", time.Since(begin),
				"err", err,
			)
		}
	}(time.Now())

	return s.next.Index(req)
}

func (s *loggingMiddleware) Error(req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		if err != nil {
			_ = s.logger.Log(
				"method", s.packageName+"_error",
				"took", time.Since(begin),
				"err", err,
			)
		}
	}(time.Now())

	return s.next.Error(req)
}

func (s *loggingMiddleware) Grpc(req interface{}) (_ interface{}, err error) {
	defer func(begin time.Time) {
		if err != nil {
			_ = s.logger.Log(
				"method", s.packageName+"_grpc",
				"took", time.Since(begin),
				"err", err,
			)
		}
	}(time.Now())

	return s.next.Grpc(req)
}
