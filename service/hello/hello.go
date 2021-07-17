package hello

import (
	"context"
	"time"

	"github.com/Zhanat87/common-libs/tracers"
	"github.com/openzipkin/zipkin-go"
)

const (
	PackageName = "hello"
	ServiceName = "hello service"
	BaseURL     = "/api/v1/hello/"
	Greeting    = "Hi, "
)

type Service interface {
	SayHi(ctx context.Context, name string) string
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) SayHi(ctx context.Context, name string) string {
	var (
		queryLabel = "GetExamplesByParam"
		query      = "select * from example where param = :value"
	)
	//traceID, spanName := utils.GetTraceIDAndSpanNameFromContext(ctx, gokitmiddlewares.TraceEndpointNamePrefix)
	//traceidint, _ := strconv.Atoi(traceID)
	//println("test", traceID, spanName, "test")
	//sc := model.SpanContext{
	//	TraceID: model.TraceID{Low: uint64(traceidint)},
	//	//ID:      spanName,
	//}
	//span := tracers.ZipkinTracer.StartSpan(queryLabel, zipkin.Parent(sc))
	//utils.PrintContextInternals("SayHi", ctx, false)
	span, _ := tracers.ZipkinTracer.StartSpanFromContext(
		ctx,
		// ServiceName,
		queryLabel,
	)
	// utils.PrintContextInternals("SayHi StartSpanFromContext", ctx2, false)
	// spanContext := span.Context()
	// fmt.Printf("span: %#v\r\n", span)
	// fmt.Printf("spanContext.TraceID: %s\r\n", spanContext.TraceID.String())
	// fmt.Printf("spanContext.ID: %s\r\n", spanContext.ID.String())
	// fmt.Printf("spanContext.ParentID: %s\r\n", spanContext.ParentID.String())
	// add interesting key/value pair to our span
	span.Tag("query", query)
	// add interesting timed event to our span
	span.Annotate(time.Now(), "query:start")
	// do the actual query...
	time.Sleep(time.Second)
	span2 := tracers.ZipkinTracer.StartSpan(queryLabel, zipkin.Parent(span.Context()))
	span2.Tag("query2", query)
	span2.Annotate(time.Now(), "query2:start")
	time.Sleep(time.Second)
	span2.Annotate(time.Now(), "query:end")
	span2.Finish()
	// let's annotate the end...
	span.Annotate(time.Now(), "query:end")
	// we're done with this span.
	span.Finish()

	return Greeting + name
}
