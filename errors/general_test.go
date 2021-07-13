/*
run tests simple as, execute goconvey from project's root directory
*/
package errors_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/Zhanat87/go-kit-hello/errors"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {
	mySetupFunction()
	retCode := m.Run()
	myTeardownFunction()
	os.Exit(retCode)
}

func mySetupFunction() {
	println("start errors package testing")
	// here we can load config data, init db connect, etc
}

func myTeardownFunction() {
	println("success end errors package testing")
	// here we can defer close db connect, etc
}

func TestArgError(t *testing.T) {
	Convey("ArgError", t, func() {
		status := 404
		errorText := "not found"
		err := &errors.ArgError{errors.ArgErrorSystemMarket, status, errorText, errorText}
		Convey("should return correct error message", func() {
			So(err.Error(), ShouldEqual, fmt.Sprintf("%d %s", status, errorText))
		})
		Convey("should set correct developer error message", func() {
			errorText = "docker-entrypoint.sh"
			err = err.SetDevMessage(errorText)
			So(err.DeveloperMessage, ShouldEqual, errorText)
		})
	})
}
