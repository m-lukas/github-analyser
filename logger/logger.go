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

//Info logs a message with the type "INFO"
func Info(message string) {
	Log(TYPE_INFO, message)
}

//Warn logs a message with the type "WARN"
func Warn(message string) {
	Log(TYPE_WARN, message)
}

//Error logs a message with the type "ERROR"
func Error(message string) {
	Log(TYPE_ERROR, message)
}

//ErrorNoMail logs a message with the type "ERROR*"
func ErrorNoMail(message string) {
	Log(TYPE_ERROR_NO_MAIL, message)
}

//Log log a message with the given type
func Log(logType string, message string) {

	var sendMail bool
	var c string

	//define asci color codes
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

	//send log message to standart output
	fmt.Printf("%s%s[%s] %s\n", reset, c, logType, message)

	//send log mail if enabled for this type
	if sendMail {
		sendLogMail(logType, message)
	}
}

//sendLogMail send an email to the log receiver with the given type as information and message
func sendLogMail(logType string, message string) {
	defaultReceiver := os.Getenv("MAILER_LOG_RECEIVER")
	backendURL := os.Getenv("BACKEND_URL")
	enviroment := os.Getenv("ENV")

	//building subject of mail
	subject := fmt.Sprintf("[%s] Log message from (%s|%s)", logType, backendURL, enviroment)

	//create default mail and push it to buffer
	mail := mailer.NewDefaultMail([]string{defaultReceiver}, subject, message)
	mailer.Push(mail)
}
