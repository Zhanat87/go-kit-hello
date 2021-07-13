package domain_test

import (
	"os"
	"testing"

	"github.com/Zhanat87/go-kit-hello/domain"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {
	mySetupFunction()
	retCode := m.Run()
	myTeardownFunction()
	os.Exit(retCode)
}

func mySetupFunction() {
	println("start domain package testing")
}

func myTeardownFunction() {
	println("success end domain package testing")
}

func TestModel(t *testing.T) {
	Convey("Model", t, func() {
		model := &domain.Model{}
		So(model, ShouldHaveSameTypeAs, &domain.Model{})
		Convey("Model SayHi", func() {
			greeting := "Hi, "
			So(greeting, ShouldEqual, model.SayHi())
			model.Name = "test name"
			So(greeting+model.Name, ShouldEqual, model.SayHi())
		})
	})
}
