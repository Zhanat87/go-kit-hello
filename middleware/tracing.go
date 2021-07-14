package middleware

import (
	"github.com/go-kit/kit/endpoint"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
)

func GetTraceEndpoint(endPoint endpoint.Endpoint, name string) endpoint.Endpoint {
	return kitoc.TraceEndpoint("gokit:endpoint hello " + name)(endPoint)
}
