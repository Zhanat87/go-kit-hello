package grpc

import (
	"context"

	"github.com/Zhanat87/go-kit-hello/contracts"
	"github.com/Zhanat87/go-kit-hello/middleware"
	"github.com/Zhanat87/go-kit-hello/transport"

	"github.com/go-kit/kit/log"
	gokittransport "github.com/go-kit/kit/transport"
	"github.com/go-kit/kit/transport/grpc"
)

type helloGrpcServer struct {
	UnimplementedHelloServiceServer
	sayHi grpc.Handler
}

func NewServer(s contracts.HTTPService, logger log.Logger) HelloServiceServer {
	options := []grpc.ServerOption{
		grpc.ServerErrorHandler(gokittransport.NewLogErrorHandler(logger)),
	}

	return &helloGrpcServer{
		sayHi: grpc.NewServer(
			middleware.MakeIndexEndpoint(s),
			decodeHelloRequest,
			encodeHelloResponse,
			options...,
		),
	}
}

func (s *helloGrpcServer) SayHi(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	_, resp, err := s.sayHi.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*HelloResponse), nil
}

func decodeHelloRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*HelloRequest)

	return transport.HelloRequest{
		Name: req.Name,
	}, nil
}

func encodeHelloResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*transport.HelloResponse)

	return &HelloResponse{Data: resp.Data}, nil
}
