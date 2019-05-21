package mailer

import (
	"fmt"
	"time"
)

var mailBuffer = make(chan *Mail, 99)

func Push(mail *Mail) {
	mailBuffer <- mail
}

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

		case <-time.After(50 * time.Millisecond):
			break
		}
	}
}
