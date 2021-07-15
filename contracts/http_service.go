package contracts

import "context"

type HTTPService interface {
	Index(req interface{}) (interface{}, error)
	Error(req interface{}) (interface{}, error)
	Grpc(ctx context.Context, req interface{}) (interface{}, error)
}
