package hello_test

import (
	"context"
	"os"
	"testing"

	"github.com/Zhanat87/common-libs/tracers"

	"github.com/Zhanat87/go-kit-hello/service/hello"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {
	mySetupFunction()
	retCode := m.Run()
	myTeardownFunction()
	os.Exit(retCode)
}

func mySetupFunction() {
	println("start service hello package testing")
	err := tracers.InitZipkinTracerAndZipkinHTTPReporter(hello.ServiceName, ":0")
	if err != nil {
		panic(err)
	}
}

func myTeardownFunction() {
	println("success end service hello package testing")
}

func TestService(t *testing.T) {
	Convey("Service", t, func() {
		service := hello.NewService()
		name := "test name"
		So(hello.Greeting+name, ShouldEqual, service.SayHi(context.Background(), name))
	})
}
