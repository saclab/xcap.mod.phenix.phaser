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
	// Logger
	logger.InitD()

	// A simple cosmetic box
	res, _ := strutil.DrawBox("XCAP Module | xcap.mod.phenix.phaser", 100, strutil.Center)
	fmt.Println(res)

	// Commandline paramaters and validation
	// Assumptions:
	// 1. If model identity is not given, we would treat the model as an homology model.
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

	// 1. Generate the configuration file for phenix.phaser
	// The configuration file is kept in the specified output directory
	// unitcell and spacegroup are read from the input sca reflection file
	runner.GeneraeteConfig(*reflectionFile, *modelFile, *modelIdentity, *outputDir)

	//2. Run the binary
	runner.RunBin(*outputDir+"/config.eff", *outputDir)
}
