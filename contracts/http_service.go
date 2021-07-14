package contracts

type HTTPService interface {
	Index(req interface{}) (interface{}, error)
	Error(req interface{}) (interface{}, error)
}
