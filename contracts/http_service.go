package contracts

type HTTPService interface {
	Index(req interface{}) (interface{}, error)
}
