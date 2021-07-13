package transport

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Data string `json:"data"`
}
