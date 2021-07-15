package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Zhanat87/common-libs/httphandlers"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/Zhanat87/go-kit-hello/factory"
	"github.com/Zhanat87/go-kit-hello/middleware"
	"github.com/Zhanat87/go-kit-hello/utils"

	commongrpc "github.com/Zhanat87/common-libs/grpc"
	"github.com/Zhanat87/common-libs/loggers"
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
)

func main() {
	httpAddr := flag.String("http.addr", ":8080", "HTTP listen address only port :8080")
	grpcAddr := flag.String("grpc.addr", ":50051", "gRPC listen address only port :50051")
	flag.Parse()
	logger := new(loggers.GoKitLoggerFactory).CreateLogger()
	httpLogger := log.With(logger, "component", "http")
	// Set-up our OpenCensus instrumentation with Zipkin backend
	zipkinURL := "http://localhost:9411/api/v2/spans"
	var tracer *zipkin.Tracer
	{
		var (
			reporter         = httpreporter.NewReporter(zipkinURL)
			localEndpoint, _ = zipkin.NewEndpoint(utils.PackageName, ":0")
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
	helloHTTPService := new(factory.ServiceFactory).CreateHTTPService(httpLogger, tracer)
	mux.Handle(utils.BaseURL, hellohttp.MakeHandler(middleware.MakeEndpoints(helloHTTPService), httpLogger,
		utils.BaseURL, hellohttp.DecodeIndexRequest))
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/health-check", httphandlers.HealthCheck)
	http.Handle("/api/v1/", httphandlers.AccessControl(mux))
	errs := make(chan error, 3)
	baseGrpcServer := grpc.NewServer(grpc.StatsHandler(zipkingrpc.NewServerHandler(tracer)))
	grpcHelloServer := hellogrpc.NewServer(helloHTTPService, logger)
	commongrpc.RegisterHelloServiceServer(baseGrpcServer, grpcHelloServer)
	reflection.Register(baseGrpcServer)
	grpcListener, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		panic(fmt.Errorf("fatal error while init gRPC listener: %s", err))
	}
	go func() {
		_ = logger.Log("transport", "http", "address", *httpAddr, "msg", "listening hello-api")
		errs <- http.ListenAndServe(*httpAddr, nil)
	}()
	go func() {
		_ = logger.Log("transport", "grpc", "address", *grpcAddr, "msg", "listening hello-api")
		errs <- baseGrpcServer.Serve(grpcListener)
	}()
	go func() {
		c := make(chan os.Signal, 2)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	_ = logger.Log("terminated", <-errs)
}
