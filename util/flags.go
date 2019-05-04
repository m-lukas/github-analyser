package util

import (
	"flag"
)

func IsTesting() bool {
	if flag.Lookup("test.v") == nil {
		return false
	}
	return true
}
