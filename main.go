package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ozgio/strutil"
	"github.com/saclab/xcap.libs/logger"
	"github.com/saclab/xcap.mod.phenix.phaser/runner"
)

// Execute a binary in go

func main() {

	logger.InitD()
	res, _ := strutil.DrawBox("XCAP Module | xcap.mod.phenix.phaser", 100, strutil.Center)
	fmt.Println(res)

	reflectionFile := flag.String("reflection", "", "--reflection pass the reflection file in sca")
	modelFile := flag.String("model", "", "--model pass the model in pdb format")
	modelIdentity := flag.Int("modelIdentity", 100, "--modelIdentity pass the model. If not passed the model is considered to be a homology model")
	outputDir := flag.String("outputDir", ".", "--outputDir point to the dir where output will be stored")

	flag.Parse()

	if *reflectionFile == "" {
		logger.Info.Println("Please specify the reflection file using paramater --reflection")
		os.Exit(2)
	}
	if *modelFile == "" {
		logger.Info.Println("Please specify the model file using paramater --model")
		os.Exit(2)
	}

	logger.Info.Println("Reflection File Path: ", *reflectionFile)
	logger.Info.Println("Model File Path: ", *modelFile)
	logger.Info.Println("Model Identity: ", *modelIdentity)
	logger.Info.Println("Output Directory: ", *outputDir)

	runner.GeneraeteConfig(*reflectionFile, *modelFile, *modelIdentity, *outputDir)

}
