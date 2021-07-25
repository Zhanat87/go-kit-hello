package ping

import "context"

const (
	PackageName = "ping"
	BaseURL     = "/api/v1/ping/"
)

type Service interface {
	Ping(ctx context.Context, url string) (string, error)
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) Ping(ctx context.Context, url string) (string, error) {
	return "", nil
}
