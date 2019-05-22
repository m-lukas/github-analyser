package mailer

import (
	"net/smtp"
	"os"
	"strings"
)

//Mail struct for building mails
type Mail struct {
	Sender   string
	Auth     smtp.Auth
	Receiver []string
	Subject  string
	Body     string
}

//DefaultSender returns the configurated mail sender
func DefaultSender() string {
	sender := os.Getenv("MAILER_USER_MAIL")
	if sender != "" {
		return sender
	} else {
		sender = "fussl.board@gmail.com"
	}
	return sender
}

//DefaultAuth returns the smtp auth using the configurated information
func DefaultAuth() smtp.Auth {
	return smtp.PlainAuth("", os.Getenv("MAILER_USER_MAIL"), os.Getenv("MAILER_USER_PASS"), os.Getenv("MAILER_SMTP_AUTH"))
}

//NewDefaultMail returns a mail object with the default sender and default auth
func NewDefaultMail(receiver []string, subject string, body string) *Mail {
	mail := &Mail{}
	mail.Sender = DefaultSender()
	mail.Auth = DefaultAuth()
	mail.Receiver = receiver
	mail.Subject = subject
	mail.Body = body
	return mail
}

//NewMail returns a fully configurateable mail object
func NewMail(sender string, auth smtp.Auth, receiver []string, subject string, body string) *Mail {
	mail := &Mail{}
	mail.Sender = sender
	mail.Auth = auth
	mail.Receiver = receiver
	mail.Subject = subject
	mail.Body = body
	return mail
}

//Send sends the mail object using smtp
func Send(mail *Mail) error {

	server := os.Getenv("MAILER_SMTP_SEND")

	message := "From: " + mail.Sender + "\n" +
		"To: " + strings.Join(mail.Receiver, ",") + "\n" +
		"Subject: " + mail.Subject + "\n\n" +
		mail.Body

	err := smtp.SendMail(server, mail.Auth, mail.Sender, mail.Receiver, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
