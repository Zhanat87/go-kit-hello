package middleware

import (
	"context"

	"github.com/Zhanat87/go-kit-hello/contracts"
	"github.com/Zhanat87/go-kit-hello/transport"
	"github.com/go-kit/kit/endpoint"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
)

type Endpoints struct {
	IndexEndpoint endpoint.Endpoint
	ErrorEndpoint endpoint.Endpoint
}

func MakeEndpoints(s contracts.HTTPService) Endpoints {
	// http://localhost:9411/zipkin/?serviceName=hello&lookback=15m&endTs=1626256523000&limit=10
	indexEndpoint := MakeIndexEndpoint(s)
	indexEndpoint = kitoc.TraceEndpoint("gokit:endpoint hello index")(indexEndpoint)

	return Endpoints{
		IndexEndpoint: indexEndpoint,
		ErrorEndpoint: MakeErrorEndpoint(s),
	}
}

func MakeIndexEndpoint(next contracts.HTTPService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.HelloRequest)
		resp, err := next.Index(&req)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}

func MakeErrorEndpoint(next contracts.HTTPService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.HelloRequest)
		resp, err := next.Error(&req)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}
