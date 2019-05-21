package logger

import (
	"fmt"
	"os"

	"github.com/m-lukas/github-analyser/mailer"
)

const (
	TYPE_ERROR         = "ERROR"
	TYPE_ERROR_NO_MAIL = "ERROR*"
	TYPE_WARN          = "WARN"
	TYPE_INFO          = "INFO"
)

func Info(message string) {
	Log(TYPE_INFO, message)
}

func Warn(message string) {
	Log(TYPE_WARN, message)
}

func Error(message string) {
	Log(TYPE_ERROR, message)
}

func ErrorNoMail(message string) {
	Log(TYPE_ERROR_NO_MAIL, message)
}

func Log(logType string, message string) {

	var sendMail bool
	var c string

	reset := "\u001b[0m"
	red := "\u001b[31;1m"
	green := "\u001b[32;1m"
	yellow := "\u001b[33;1m"
	white := "\u001b[37;1m"

	switch logType {
	case TYPE_INFO:
		c = green
		sendMail = false
	case TYPE_WARN:
		c = yellow
		sendMail = false
	case TYPE_ERROR:
		c = red
		sendMail = true
	case TYPE_ERROR_NO_MAIL:
		c = red
		sendMail = false
	default:
		c = white
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
