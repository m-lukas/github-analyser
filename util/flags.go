package util

import (
	"fmt"
	"os"
	"strings"

	"github.com/m-lukas/github-analyser/logger"
)

func IsTesting() bool {
	if strings.HasSuffix(os.Args[0], ".test") {
		return true
	}
	return false
}

func ReadBoolFlag(flagName string) bool {
	flag := os.Getenv(flagName)

	switch flag {
	case "0":
		logger.Warn(fmt.Sprintf("%s IS DISABLED.", flagName))
		return false
	case "1":
		logger.Warn(fmt.Sprintf("%s IS ENABLED.", flagName))
		return true
	default:
		logger.Warn(fmt.Sprintf("%s IS DISABLED.", flagName))
		return false
	}
}
