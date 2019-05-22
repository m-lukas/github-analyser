package errorutil

import "fmt"

const (
	InternalServerError = "Internal server error!"
)

//ConversionError is used in http-requests with invalid param types
type ConversionError struct {
	Err   error
	Param string
	Value string
}

func (e ConversionError) Error() string {
	prefix := "Conversion Error:"
	return fmt.Sprintf("%s cannot use value: %s for parameter %s!", prefix, e.Value, e.Param)
}

//NullOrEmptyError is used in http-request with invalid null or empty params
type NullOrEmptyError struct {
	Err   error
	Param string
}

func (e NullOrEmptyError) Error() string {
	prefix := "NullOrEmpty Error:"
	return fmt.Sprintf("%s %s must not be null or empty", prefix, e.Param)
}
