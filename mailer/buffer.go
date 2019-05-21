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
	var err error
	fmt.Println("Started mail worker!")
	for {
		select {
		case mail := <-mailBuffer:

			err = Send(mail)
			if err != nil {
				//Log error
				fmt.Println(err)
			} else {
				fmt.Println("Send mail!")
			}

		case <-time.After(50 * time.Millisecond):
			break
		}
	}
}
