package middleware

import (
	"context"

	"github.com/Zhanat87/go-kit-hello/contracts"
	"github.com/Zhanat87/go-kit-hello/transport"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	IndexEndpoint endpoint.Endpoint
}

func MakeEndpoints(s contracts.HTTPService) Endpoints {
	return Endpoints{
		IndexEndpoint: MakeIndexEndpoint(s),
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
