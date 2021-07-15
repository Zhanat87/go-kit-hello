package tracers

import (
	oczipkin "contrib.go.opencensus.io/exporter/zipkin"
	"github.com/Zhanat87/go-kit-hello/utils"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/model"
	httpreporter "github.com/openzipkin/zipkin-go/reporter/http"
	"go.opencensus.io/trace"
	"strconv"
)

const endpointURL = "http://localhost:9411/api/v2/spans"

func NewZipkinTracer(port string) (*zipkin.Tracer, error) {
	portInt, _ := strconv.Atoi(port)
	portUint16 := uint16(portInt)
	// The reporter sends traces to zipkin server
	reporter := httpreporter.NewReporter(endpointURL)
	// Local endpoint represent the local service information
	localEndpoint := &model.Endpoint{ServiceName: utils.PackageName, Port: portUint16}
	// Sampler tells you which traces are going to be sampled or not. In this case we will record 100% (1.00) of traces.
	sampler, err := zipkin.NewCountingSampler(1)
	if err != nil {
		return nil, err
	}
	t, err := zipkin.NewTracer(
		reporter,
		zipkin.WithSampler(sampler),
		zipkin.WithLocalEndpoint(localEndpoint),
	)
	if err != nil {
		return nil, err
	}

	return t, err
}

func NewZipkinTracer2() (*zipkin.Tracer, error) {
	// Set-up our OpenCensus instrumentation with Zipkin backend
	reporter := httpreporter.NewReporter(endpointURL)
	localEndpoint, _ := zipkin.NewEndpoint(utils.PackageName, ":0")
	exporter := oczipkin.NewExporter(reporter, localEndpoint)
	defer reporter.Close()
	// Always sample our traces for this demo.
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	// Register our trace exporter.
	trace.RegisterExporter(exporter)

	return zipkin.NewTracer(reporter)
}
