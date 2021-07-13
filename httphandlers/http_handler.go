package httphandlers

import "net/http"

type HTTPHandler struct {
	Path    string
	Handler http.Handler
	Methods []string
}
