package util

import (
	"fmt"
	"log"
	"os"
	"strings"
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
		log.Println(fmt.Sprintf("%s IS DISABLED.", flagName))
		return false
	case "1":
		log.Println(fmt.Sprintf("%s IS ENABLED.", flagName))
		return true
	default:
		log.Println(fmt.Sprintf("%s IS DISABLED.", flagName))
		return false
	}
}
