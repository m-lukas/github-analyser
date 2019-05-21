package main

import (
	"fmt"
	"log"

	"github.com/m-lukas/github-analyser/logger"
	"github.com/m-lukas/github-analyser/metrix"
	"github.com/m-lukas/github-analyser/setup"
)

func setupInit() {
	err := setup.SetupInputFile("m-lukas")
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to setup input files with error: %s", err.Error()))
		log.Fatal(err)
	}
	logger.Info("Setup for metrix was successful!")
}

func metrixInit() {
	err := metrix.CalcScoreParams()
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to calculate new scores with error: %s", err.Error()))
		log.Fatal(err)
	}
	logger.Info("Metrix initialization was successful!")
}
