package transport

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Data string `json:"data"`
}

type ErrorRequest struct {
	Text string `json:"text"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type PingRequest struct {
	Ping string `json:"ping"`
}

type PingResponse struct {
	Pong string `json:"pong"`
}
