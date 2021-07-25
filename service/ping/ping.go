package ping

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	tracingtransport "github.com/Zhanat87/go-kit-tracing/transport"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
)

const (
	PackageName = "ping"
	BaseURL     = "/api/v1/ping/"
)

type Service interface {
	Ping(ctx context.Context, url string) (string, error)
}

type service struct {
	zipkinTracer *zipkin.Tracer
}

func NewService(zipkinTracer *zipkin.Tracer) Service {
	return &service{zipkinTracer}
}

func (s *service) Ping(ctx context.Context, url string) (string, error) {
	span, ctx := s.zipkinTracer.StartSpanFromContext(ctx, "ping service by url: "+url)
	defer span.Finish()
	pongRequest := &tracingtransport.PongRequest{
		Data: "ping from hello service",
	}
	pongRequestJSON, err := json.Marshal(pongRequest)
	if err != nil {
		return "", err
	}
	newRequest, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(pongRequestJSON))
	if err != nil {
		return "", err
	}
	newRequest.Header.Set("Content-Type", "application/json")
	client, err := zipkinhttp.NewClient(s.zipkinTracer, zipkinhttp.ClientTrace(true))
	if err != nil {
		return "", err
	}
	resp, err := client.DoWithAppSpan(newRequest, "ping request from hello service")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var pongResponse tracingtransport.PongResponse
	err = json.Unmarshal(body, &pongResponse)
	if err != nil {
		return "", err
	}

	return pongResponse.Data, nil
}
