package mailer

import (
	"fmt"
	"time"
)

//mailBuffer as a queue for mails
var mailBuffer = make(chan *Mail, 99)

//Push adds the mail to the buffer/queue
func Push(mail *Mail) {
	mailBuffer <- mail
}

//StartWorker starts a worker that checks the mail buffer and sends mails if available
func StartWorker() {
	fmt.Println("\u001b[1m\u001b[32;1m[INFO] Started mail worker!")
	var err error
	for {
		select {
		case mail := <-mailBuffer:

			err = Send(mail)
			if err != nil {
				fmt.Println("\u001b[1m\u001b[31;1m[ERROR] Error occured while sending mail!")
			} else {
				fmt.Println("\u001b[1m\u001b[32;1m[INFO] Send mail!")
			}

		//slow down loop
		case <-time.After(50 * time.Millisecond):
			break
		}
	}
}
