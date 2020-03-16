package main

import (
	"fmt"
	"io/ioutil"

	"github.com/ozgio/strutil"
	"github.com/saclab/xcap.libs/execuator"
	"github.com/tidwall/gjson"
)

// Read the json file
func readFile(path string) string {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(dat)
}

// Execute a binary in go

func main() {
	res, _ := strutil.DrawBox("XCAP Module | xcap.mod.phenix.phaser", 90, strutil.Center)
	fmt.Println(res)

	data := readFile("module.json")
	binaryPath := gjson.Get(data, "binary.location").String()
	fmt.Println(binaryPath)

	execuator.Shellout(binaryPath, "--version")
}
