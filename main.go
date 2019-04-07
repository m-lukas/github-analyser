package main

import "log"

func main() {
	client, err := InitMongo()
	if err != nil {
		log.Fatal(err)
	}
}
