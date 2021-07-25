package middleware

import (
	"context"

	"github.com/Zhanat87/common-libs/gokitmiddlewares"
	errorservice "github.com/Zhanat87/go-kit-hello/service/error"
	"github.com/Zhanat87/go-kit-hello/service/hello"
	"github.com/Zhanat87/go-kit-hello/service/ping"
	"github.com/openzipkin/zipkin-go"
)

type helloZipkinTracingMiddleware struct {
	next         hello.HTTPService
	zipkinTracer gokitmiddlewares.Tracer
}

func NewHelloZipkinTracingMiddleware(s hello.HTTPService,
	packageName string, zipkinTracer *zipkin.Tracer) hello.HTTPService {
	return &helloZipkinTracingMiddleware{
		next:         s,
		zipkinTracer: gokitmiddlewares.NewZipkinTracing(zipkinTracer, packageName),
	}
}

func (s *helloZipkinTracingMiddleware) Index(ctx context.Context, req interface{}) (_ interface{}, err error) {
	span, ctx := s.zipkinTracer.Trace(ctx, "index")
	defer span.Finish()

	return s.next.Index(ctx, req)
}

type errorZipkinTracingMiddleware struct {
	next         errorservice.HTTPService
	zipkinTracer gokitmiddlewares.Tracer
}

func NewErrorZipkinTracingMiddleware(s errorservice.HTTPService,
	packageName string, zipkinTracer *zipkin.Tracer) errorservice.HTTPService {
	return &errorZipkinTracingMiddleware{
		next:         s,
		zipkinTracer: gokitmiddlewares.NewZipkinTracing(zipkinTracer, packageName),
	}
}

func (s *errorZipkinTracingMiddleware) Index(ctx context.Context, req interface{}) (_ interface{}, err error) {
	span, ctx := s.zipkinTracer.Trace(ctx, "index")
	defer span.Finish()

	return s.next.Index(ctx, req)
}

type pingZipkinTracingMiddleware struct {
	next         ping.HTTPService
	zipkinTracer gokitmiddlewares.Tracer
}

func NewPingZipkinTracingMiddleware(s ping.HTTPService,
	packageName string, zipkinTracer *zipkin.Tracer) ping.HTTPService {
	return &pingZipkinTracingMiddleware{
		next:         s,
		zipkinTracer: gokitmiddlewares.NewZipkinTracing(zipkinTracer, packageName),
	}
}

func (s *pingZipkinTracingMiddleware) HTTP(ctx context.Context, req interface{}) (_ interface{}, err error) {
	span, ctx := s.zipkinTracer.Trace(ctx, "http")
	defer span.Finish()

	return s.next.HTTP(ctx, req)
}

func (s *pingZipkinTracingMiddleware) Grpc(ctx context.Context, req interface{}) (_ interface{}, err error) {
	span, ctx := s.zipkinTracer.Trace(ctx, "grpc")
	defer span.Finish()

	return s.next.Grpc(ctx, req)
}
