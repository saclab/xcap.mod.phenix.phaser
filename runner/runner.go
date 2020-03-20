package runner

import (
	"fmt"
	"os"
	"text/template"

	"github.com/saclab/xcap.libs/executor"
	"github.com/saclab/xcap.libs/logger"
	"github.com/saclab/xcap.mod.phenix.phaser/lib"
	"github.com/saclab/xcap.mod.phenix.phaser/templates/phenixPhaserInputConfig"
	"github.com/tidwall/gjson"
)

// GeneraeteConfig : Generate the configuration file in the output dir
func GeneraeteConfig(reflectionFile string, modelFile string, modelIdentity int, outputDir string) {
	// Logger deamon
	logger.InitD()

	// 1. check if output dir exist if not create it
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		logger.Warning.Println("Output Dir does not exist, will create -> ", outputDir)
		err := os.MkdirAll(outputDir, os.FileMode(0750))
		if err != nil {
			logger.Error.Println("Failed to create dir : ", outputDir, err)
			os.Exit(2)
		}
	}
	logger.Info.Println("Using Output Dir : ", outputDir)

	// 2. create the template blank file
	// We wil place the blank file in the output dir
	outputFile, err := os.Create(outputDir + "/config.eff")
	if err != nil {
		logger.Error.Println("Failed to create config file : ", outputDir, err)
		os.Exit(2)
	}

	// 3. make a key value map for the template
	vars := make(map[string]interface{})
	vars["reflectionFile"] = reflectionFile
	vars["modelFile"] = modelFile
	vars["modelIdentity"] = modelIdentity
	vars["outputDir"] = outputDir

	// 4. detect the space group and unit cell
	vars["spacegroup"], vars["unitcell"] = lib.DetectSpaceGroupAndUnitCellFromSCA(reflectionFile)

	// 5. Parse the template and execute
	// The template is at templates/phenixPhaserInputConfig/phenixPhaserInputConfig.go
	t := template.Must(template.New("letter").Parse(phenixPhaserInputConfig.Tmpl))
	err = t.Execute(outputFile, vars)
	if err != nil {
		logger.Error.Println("Failed to create config file (Template) : ", outputDir, err)
		os.Exit(2)
	}
	logger.Info.Println("Template created: ", outputDir+"/config.eff")
}

// RunBin : Run the binary function
func RunBin() {

	config := lib.ReadFile("module.json")
	binaryPath := gjson.Get(config, "binary.location").String()
	fmt.Println(binaryPath)
	executor.Shellout(binaryPath, "dev/example.eff")

}
