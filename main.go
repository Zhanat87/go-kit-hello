package main

import (
	"flag"
	"fmt"
	"github.com/Zhanat87/go-kit-hello/httphandlers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Zhanat87/go-kit-hello/factory"
	"github.com/Zhanat87/go-kit-hello/middleware"
	"github.com/Zhanat87/go-kit-hello/utils"

	"github.com/Zhanat87/go-kit-hello/loggers"
	hellohttp "github.com/Zhanat87/go-kit-hello/transport/http"
	"github.com/go-kit/kit/log"
)

func main() {
	httpAddr := flag.String("http.addr", ":8080", "HTTP listen address only port :8080")
	flag.Parse()
	logger := new(loggers.GoKitLoggerFactory).CreateLogger()
	httpLogger := log.With(logger, "component", "http")
	mux := http.NewServeMux()
	// curl -i -X POST -H "Content-Type: application/json" -d '{"name":"val"}' http://localhost:8080/api/v1/hello
	mux.Handle(utils.BaseURL, hellohttp.MakeHandler(middleware.MakeEndpoints(
		new(factory.ServiceFactory).CreateHTTPService(httpLogger)), httpLogger,
		utils.BaseURL, hellohttp.DecodeIndexRequest))
	// http://localhost:8080/metrics
	http.Handle("/metrics", promhttp.Handler())
	// http://localhost:8080/check
	http.HandleFunc("/check", httphandlers.HealthCheck)
	http.Handle("/api/v1/", httphandlers.AccessControl(mux))
	errs := make(chan error, 2)
	go func() {
		_ = logger.Log("transport", "http", "address", *httpAddr, "msg", "listening hello-api")
		errs <- http.ListenAndServe(*httpAddr, nil)
	}()
	go func() {
		c := make(chan os.Signal, 2)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	_ = logger.Log("terminated", <-errs)
}
