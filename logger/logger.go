package logger

import (
	"fmt"
	"os"

	"github.com/m-lukas/github-analyser/mailer"
)

const (
	ERROR = "ERROR"
	WARN  = "WARN"
	INFO  = "INFO"
)

func Log(logType string, message string) {

	var sendMail bool
	var c string

	reset := "\u001b[0m"
	red := "\u001b[1m\u001b[31;1m"
	green := "\u001b[1m\u001b[32;1m"
	yellow := "\u001b[1m\u001b[33;1m"

	switch logType {
	case WARN:
		c = yellow
		sendMail = false
	case ERROR:
		c = red
		sendMail = true
	default:
		c = green
		sendMail = false
	}

	fmt.Printf("%s%s[%s] %s\n", reset, c, logType, message)

	if sendMail {
		sendLogMail(logType, message)
	}
}

func sendLogMail(logType string, message string) {
	defaultReceiver := os.Getenv("MAILER_LOG_RECEIVER")
	backendURL := os.Getenv("BACKEND_URL")
	enviroment := os.Getenv("ENV")

	subject := fmt.Sprintf("[%s] Log message from (%s|%s)", logType, backendURL, enviroment)

	mail := mailer.NewDefaultMail([]string{defaultReceiver}, subject, message)
	mailer.Push(mail)
}
