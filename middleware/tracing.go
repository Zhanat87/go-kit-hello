package middleware

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/model"
)

func GetTraceEndpoint(endPoint endpoint.Endpoint, name string) endpoint.Endpoint {
	return kitoc.TraceEndpoint("gokit:endpoint hello " + name)(endPoint)
}

func TraceEndpoint(tracer *zipkin.Tracer, name string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			var sc model.SpanContext
			if parentSpan := zipkin.SpanFromContext(ctx); parentSpan != nil {
				sc = parentSpan.Context()
			}
			sp := tracer.StartSpan(name, zipkin.Parent(sc))
			defer sp.Finish()
			ctx = zipkin.NewContext(ctx, sp)

			return next(ctx, request)
		}
	}
}
