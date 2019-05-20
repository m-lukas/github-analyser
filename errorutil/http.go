package errorutil

import "fmt"

const (
	InternalServerError = "Internal server error!"
)

type ConversionError struct {
	Err   error
	Param string
	Value string
}

func (e ConversionError) Error() string {
	prefix := "Conversion Error:"
	return fmt.Sprintf("%s cannot use value: %s for parameter %s!", prefix, e.Value, e.Param)
}

type NullOrEmptyError struct {
	Err   error
	Param string
}

func (e NullOrEmptyError) Error() string {
	prefix := "NullOrEmpty Error:"
	return fmt.Sprintf("%s %s must not be null or empty", prefix, e.Param)
}
