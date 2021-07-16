package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	errorservice "github.com/Zhanat87/go-kit-hello/service/error"

	commongrpc "github.com/Zhanat87/common-libs/grpc"
	"github.com/Zhanat87/common-libs/httphandlers"
	"github.com/Zhanat87/common-libs/loggers"
	"github.com/Zhanat87/go-kit-hello/factory"
	"github.com/Zhanat87/go-kit-hello/middleware"
	"github.com/Zhanat87/go-kit-hello/service/hello"
	hellogrpc "github.com/Zhanat87/go-kit-hello/transport/grpc"
	hellohttp "github.com/Zhanat87/go-kit-hello/transport/http"
	"github.com/go-kit/kit/log"

	oczipkin "contrib.go.opencensus.io/exporter/zipkin"
	zipkin "github.com/openzipkin/zipkin-go"
	zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	httpreporter "github.com/openzipkin/zipkin-go/reporter/http"
	"go.opencensus.io/trace"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	httpAddr := os.Getenv("HTTP_ADDR")
	grpcAddr := os.Getenv("GRPC_ADDR")
	logger := new(loggers.GoKitLoggerFactory).CreateLogger()
	httpLogger := log.With(logger, "component", "http")
	// todo: вынести отсюда
	// Set-up our OpenCensus instrumentation with Zipkin backend
	zipkinURL := "http://localhost:9411/api/v2/spans"
	var tracer *zipkin.Tracer
	{
		var (
			reporter         = httpreporter.NewReporter(zipkinURL)
			localEndpoint, _ = zipkin.NewEndpoint(hello.PackageName, ":0")
			exporter         = oczipkin.NewExporter(reporter, localEndpoint)
			err              error
		)
		defer reporter.Close()
		// Always sample our traces for this demo.
		trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
		// Register our trace exporter.
		trace.RegisterExporter(exporter)
		tracer, err = zipkin.NewTracer(reporter)
		if err != nil {
			panic(err)
		}
	}
	mux := http.NewServeMux()
	helloHTTPService := new(factory.HelloServiceFactory).CreateHTTPService(httpLogger)
	mux.Handle(hello.BaseURL, hellohttp.MakeHelloHandler(
		middleware.MakeHelloEndpoints(helloHTTPService), httpLogger, hello.BaseURL))
	mux.Handle(errorservice.BaseURL, hellohttp.MakeErrorHandler(
		middleware.MakeErrorEndpoints(new(factory.ErrorServiceFactory).CreateHTTPService(httpLogger)),
		httpLogger, errorservice.BaseURL))
	httphandlers.InitDefaultHandlers(mux)
	errs := make(chan error, 3)
	baseGrpcServer := grpc.NewServer(grpc.StatsHandler(zipkingrpc.NewServerHandler(tracer)))
	grpcHelloServer := hellogrpc.NewServer(helloHTTPService, logger)
	commongrpc.RegisterHelloServiceServer(baseGrpcServer, grpcHelloServer)
	reflection.Register(baseGrpcServer)
	grpcListener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		panic(fmt.Errorf("fatal error while init gRPC listener: %s", err))
	}
	go func() {
		_ = logger.Log("transport", "http", "address", httpAddr, "msg", "listening hello-api")
		errs <- http.ListenAndServe(httpAddr, nil)
	}()
	go func() {
		_ = logger.Log("transport", "grpc", "address", grpcAddr, "msg", "listening hello-api")
		errs <- baseGrpcServer.Serve(grpcListener)
	}()
	go func() {
		c := make(chan os.Signal, 2)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	_ = logger.Log("terminated", <-errs)
}

// сделать пинг сервис через grpc здесь и вынести в новый сервис понг
// сделать пинг сервис через http здесь и вынести в новый сервис понг
// сделать zipkin nested span и проверить как это все работает
