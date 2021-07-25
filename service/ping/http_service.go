package ping

import (
	"context"
	"fmt"

	commongrpc "github.com/Zhanat87/common-libs/grpc"
	"github.com/Zhanat87/go-kit-hello/transport"
	"github.com/openzipkin/zipkin-go"
	zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	"google.golang.org/grpc"
)

type HTTPService interface {
	Grpc(ctx context.Context, req interface{}) (response interface{}, err error)
	HTTP(ctx context.Context, req interface{}) (response interface{}, err error)
}

type httpService struct {
	service      Service
	zipkinTracer *zipkin.Tracer
}

func NewHTTPService(zipkinTracer *zipkin.Tracer) HTTPService {
	return &httpService{service: NewService(zipkinTracer), zipkinTracer: zipkinTracer}
}

func (s *httpService) Grpc(ctx context.Context, req interface{}) (interface{}, error) {
	span, ctx := s.zipkinTracer.StartSpanFromContext(ctx, "ping grpc service span")
	defer span.Finish()
	pingRequest, ok := req.(*transport.PingRequest)
	if !ok {
		return nil, fmt.Errorf("error convert transport.PingRequest: %#v", req)
	}
	connection, err := grpc.Dial(pingRequest.Ping,
		grpc.WithInsecure(),
		grpc.WithStatsHandler(zipkingrpc.NewClientHandler(s.zipkinTracer)),
	)
	if err != nil {
		return nil, err
	}
	defer connection.Close()
	client := commongrpc.NewHelloServiceClient(connection)
	helloRequest := &commongrpc.HelloRequest{Name: "test grpc name"}
	helloResponse, err := client.SayHi(ctx, helloRequest)
	if err != nil {
		return nil, err
	}

	return &transport.PingRequest{Ping: helloResponse.Data}, nil
}

func (s *httpService) HTTP(ctx context.Context, req interface{}) (interface{}, error) {
	span, ctx := s.zipkinTracer.StartSpanFromContext(ctx, "ping http service span")
	defer span.Finish()
	pingRequest, ok := req.(*transport.PingRequest)
	if !ok {
		return nil, fmt.Errorf("error convert transport.PingRequest: %#v", req)
	}
	pong, err := s.service.Ping(ctx, pingRequest.Ping)
	if err != nil {
		return nil, err
	}

	return &transport.PingResponse{Pong: pong}, nil
}
