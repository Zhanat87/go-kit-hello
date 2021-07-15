package middleware

import (
	"context"

	"github.com/Zhanat87/go-kit-hello/service/hello"
	"github.com/Zhanat87/go-kit-hello/transport"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	IndexEndpoint endpoint.Endpoint
	ErrorEndpoint endpoint.Endpoint
	GrpcEndpoint  endpoint.Endpoint
}

func MakeEndpoints(s hello.HTTPService) Endpoints {
	// http://localhost:9411/zipkin/?serviceName=hello&lookback=15m&endTs=1626256523000&limit=10
	return Endpoints{
		IndexEndpoint: GetTraceEndpoint(MakeIndexEndpoint(s), "index"),
		ErrorEndpoint: GetTraceEndpoint(MakeErrorEndpoint(s), "error"),
		GrpcEndpoint:  GetTraceEndpoint(MakeGrpcEndpoint(s), "grpc"),
	}
}

func MakeIndexEndpoint(next hello.HTTPService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.HelloRequest)
		resp, err := next.Index(&req)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}

func MakeErrorEndpoint(next hello.HTTPService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.HelloRequest)
		resp, err := next.Error(&req)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}

func MakeGrpcEndpoint(next hello.HTTPService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.HelloRequest)
		resp, err := next.Grpc(ctx, &req)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}
