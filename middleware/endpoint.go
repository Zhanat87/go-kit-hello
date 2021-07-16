package middleware

import (
	"context"

	"github.com/Zhanat87/common-libs/gokitmiddlewares"
	errorservice "github.com/Zhanat87/go-kit-hello/service/error"
	"github.com/Zhanat87/go-kit-hello/service/hello"
	"github.com/Zhanat87/go-kit-hello/transport"
	"github.com/go-kit/kit/endpoint"
)

type HelloEndpoints struct {
	IndexEndpoint endpoint.Endpoint
}

func MakeHelloEndpoints(s hello.HTTPService) HelloEndpoints {
	return HelloEndpoints{
		IndexEndpoint: gokitmiddlewares.GetTraceEndpoint(MakeHelloIndexEndpoint(s), "hello index"),
	}
}

func MakeHelloIndexEndpoint(next hello.HTTPService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.HelloRequest)
		resp, err := next.Index(&req)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}

type ErrorEndpoints struct {
	IndexEndpoint endpoint.Endpoint
}

func MakeErrorEndpoints(s errorservice.HTTPService) ErrorEndpoints {
	return ErrorEndpoints{
		IndexEndpoint: gokitmiddlewares.GetTraceEndpoint(MakeErrorIndexEndpoint(s), "error index"),
	}
}

func MakeErrorIndexEndpoint(next errorservice.HTTPService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.ErrorRequest)
		resp, err := next.Index(&req)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}

//type PingEndpoints struct {
//	GrpcEndpoint endpoint.Endpoint
//	//HTTPEndpoint endpoint.Endpoint
//}
//
//func MakePingEndpoints(s hello.HTTPService) PingEndpoints {
//	return PingEndpoints{
//		GrpcEndpoint: GetTraceEndpoint(MakePingGrpcEndpoint(s), "ping grpc"),
//	}
//}
//
//func MakePingGrpcEndpoint(next hello.HTTPService) endpoint.Endpoint {
//	return func(ctx context.Context, request interface{}) (interface{}, error) {
//		req := request.(transport.HelloRequest)
//		resp, err := next.Grpc(ctx, &req)
//		if err != nil {
//			return nil, err
//		}
//
//		return resp, nil
//	}
//}
