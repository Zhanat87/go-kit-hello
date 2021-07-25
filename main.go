package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	commongrpc "github.com/Zhanat87/common-libs/grpc"
	"github.com/Zhanat87/common-libs/httphandlers"
	"github.com/Zhanat87/common-libs/loggers"
	"github.com/Zhanat87/common-libs/tracers"
	"github.com/Zhanat87/go-kit-hello/factory"
	"github.com/Zhanat87/go-kit-hello/middleware"
	errorservice "github.com/Zhanat87/go-kit-hello/service/error"
	"github.com/Zhanat87/go-kit-hello/service/hello"
	"github.com/Zhanat87/go-kit-hello/service/ping"
	appgrpc "github.com/Zhanat87/go-kit-hello/transport/grpc"
	apphttp "github.com/Zhanat87/go-kit-hello/transport/http"
	"github.com/go-kit/kit/log"

	zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	httpAddr := os.Getenv("HTTP_ADDR")
	grpcAddr := os.Getenv("GRPC_ADDR")
	logger := new(loggers.GoKitLoggerFactory).CreateLogger()
	httpLogger := log.With(logger, "component", "http")
	serviceName := os.Getenv("SERVICE_NAME")
	err := tracers.InitZipkinTracerAndZipkinHTTPReporter(serviceName, ":0")
	if err != nil {
		panic(err)
	}
	defer tracers.ZipkinReporter.Close()
	mux := http.NewServeMux()
	helloHTTPService := new(factory.HelloServiceFactory).CreateHTTPService(hello.PackageName, httpLogger, tracers.ZipkinTracer)
	mux.Handle(hello.BaseURL, apphttp.MakeHelloHandler(
		middleware.MakeHelloEndpoints(helloHTTPService), httpLogger, hello.BaseURL))
	mux.Handle(errorservice.BaseURL, apphttp.MakeErrorHandler(
		middleware.MakeErrorEndpoints(new(factory.ErrorServiceFactory).CreateHTTPService(errorservice.PackageName, httpLogger, tracers.ZipkinTracer)),
		httpLogger, errorservice.BaseURL))
	mux.Handle(ping.BaseURL, apphttp.MakePingHandler(
		middleware.MakePingEndpoints(new(factory.PingServiceFactory).CreateHTTPService(ping.PackageName, httpLogger, tracers.ZipkinTracer)),
		httpLogger, ping.BaseURL))
	httphandlers.InitDefaultHandlers(mux)
	errs := make(chan error, 3)
	baseGrpcServer := grpc.NewServer(grpc.StatsHandler(zipkingrpc.NewServerHandler(tracers.ZipkinTracer)))
	grpcHelloServer := appgrpc.NewServer(helloHTTPService, logger)
	commongrpc.RegisterHelloServiceServer(baseGrpcServer, grpcHelloServer)
	reflection.Register(baseGrpcServer)
	grpcListener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		panic(fmt.Errorf("fatal error while init gRPC listener: %s", err))
	}
	go func() {
		_ = logger.Log("transport", "http", "address", httpAddr, "msg", "listening "+serviceName+" api")
		errs <- http.ListenAndServe(httpAddr, nil)
	}()
	go func() {
		_ = logger.Log("transport", "grpc", "address", grpcAddr, "msg", "listening "+serviceName+" api")
		errs <- baseGrpcServer.Serve(grpcListener)
	}()
	go func() {
		c := make(chan os.Signal, 2)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	_ = logger.Log("terminated", <-errs)
}
