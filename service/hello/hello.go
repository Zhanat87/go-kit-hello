package hello

const (
	PackageName = "hello"
	BaseURL     = "/api/v1/hello/"
	Greeting    = "Hi, "
)

type Service interface {
	SayHi(name string) string
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) SayHi(name string) string {
	return Greeting + name
}
