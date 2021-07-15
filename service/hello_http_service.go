package service

import (
	"context"
	"errors"
	"github.com/openzipkin/zipkin-go"
	"google.golang.org/grpc"
	"log"

	"github.com/Zhanat87/go-kit-hello/contracts"
	"github.com/Zhanat87/go-kit-hello/transport"
	hellogrpc "github.com/Zhanat87/go-kit-hello/transport/grpc"
	zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
)

type helloHTTPService struct {
	helloService contracts.HelloService
	tracer       *zipkin.Tracer
}

func NewHTTPService(tracer *zipkin.Tracer) contracts.HTTPService {
	return &helloHTTPService{helloService: NewHelloService(), tracer: tracer}
}

func (s *helloHTTPService) Index(req interface{}) (interface{}, error) {
	r := req.(*transport.HelloRequest)

	return &transport.HelloResponse{Data: s.helloService.SayHi(r.Name)}, nil
}

func (s *helloHTTPService) Error(req interface{}) (interface{}, error) {
	return &transport.HelloResponse{Data: "error response"}, errors.New("error from hello")
}

func (s *helloHTTPService) Grpc(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*transport.HelloRequest)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":50051", grpc.WithInsecure(),
		grpc.WithStatsHandler(zipkingrpc.NewClientHandler(s.tracer)))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := hellogrpc.NewHelloServiceClient(conn)
	response, err := c.SayHi(context.Background(), &hellogrpc.HelloRequest{Name: r.Name})
	if err != nil {
		return nil, err
	}

	return &transport.HelloResponse{Data: response.Data}, nil
}
