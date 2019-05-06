package util

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func IsTesting() bool {
	if flag.Lookup("test.v") == nil {
		return false
	}
	return true
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
