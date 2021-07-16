package hello

import (
	"github.com/Zhanat87/go-kit-hello/transport"
)

type HTTPService interface {
	Index(req interface{}) (response interface{}, err error)
}

type httpService struct {
	service Service
}

func NewHTTPService() HTTPService {
	return &httpService{service: NewService()}
}

func (s *httpService) Index(req interface{}) (interface{}, error) {
	r := req.(*transport.HelloRequest)

	return &transport.HelloResponse{Data: s.service.SayHi(r.Name)}, nil
}

//func (s *httpService) Error(req interface{}) (interface{}, error) {
//	return &transport.HelloResponse{Data: "error response"}, errors.New("error from hello")
//}
//
//func (s *httpService) Grpc(ctx context.Context, req interface{}) (interface{}, error) {
//	r := req.(*transport.HelloRequest)
//	var conn *grpc.ClientConn
//	conn, err := grpc.Dial(":50051", grpc.WithInsecure(),
//		grpc.WithStatsHandler(zipkingrpc.NewClientHandler(s.tracer)))
//	if err != nil {
//		log.Fatalf("did not connect: %s", err)
//	}
//	defer conn.Close()
//	c := commongrpc.NewHelloServiceClient(conn)
//	response, err := c.SayHi(context.Background(), &commongrpc.HelloRequest{Name: r.Name})
//	if err != nil {
//		return nil, err
//	}
//
//	return &transport.HelloResponse{Data: response.Data}, nil
//}
