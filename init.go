package main

import (
	"fmt"
	"log"

	"github.com/m-lukas/github-analyser/metrix"
	"github.com/m-lukas/github-analyser/setup"
)

func setupInit() {
	err := setup.SetupInputFile("jo-fr")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Setup successful!")
}

func metrixInit() {
	err := metrix.CalcScoreParams()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Metrix successful!")
}
