package errors

import "fmt"

const (
	ArgErrorSystemMarket = "MARKET"
)

/////////////////////////////////
// Error structure & methods
/////////////////////////////////
type ArgError struct {
	System           string `json:"system"`
	Status           int    `json:"status"`
	Message          string `json:"message"`
	DeveloperMessage string `json:"developerMessage"`
}

func (e *ArgError) Error() string {
	return fmt.Sprintf("%d %s", e.Status, e.DeveloperMessage)
}

func (e *ArgError) SetDevMessage(developMessage string) *ArgError {
	e.DeveloperMessage = developMessage

	return e
}
