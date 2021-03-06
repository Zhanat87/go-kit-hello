package grpc

import (
	"context"

	commongrpc "github.com/Zhanat87/common-libs/grpc"
	"github.com/Zhanat87/go-kit-hello/middleware"
	"github.com/Zhanat87/go-kit-hello/service/hello"
	"github.com/Zhanat87/go-kit-hello/transport"
	"github.com/go-kit/kit/log"
	gokittransport "github.com/go-kit/kit/transport"
	"github.com/go-kit/kit/transport/grpc"
)

type helloGrpcServer struct {
	commongrpc.UnimplementedHelloServiceServer
	sayHi grpc.Handler
}

func NewServer(s hello.HTTPService, logger log.Logger) commongrpc.HelloServiceServer {
	options := []grpc.ServerOption{
		grpc.ServerErrorHandler(gokittransport.NewLogErrorHandler(logger)),
	}

	return &helloGrpcServer{
		sayHi: grpc.NewServer(
			middleware.MakeHelloIndexEndpoint(s),
			decodeHelloRequest,
			encodeHelloResponse,
			options...,
		),
	}
}

func (s *helloGrpcServer) SayHi(ctx context.Context, req *commongrpc.HelloRequest) (*commongrpc.HelloResponse, error) {
	_, resp, err := s.sayHi.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*commongrpc.HelloResponse), nil
}

func decodeHelloRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*commongrpc.HelloRequest)

	return transport.HelloRequest{
		Name: req.Name,
	}, nil
}

func encodeHelloResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*transport.HelloResponse)

	return &commongrpc.HelloResponse{Data: resp.Data}, nil
}
