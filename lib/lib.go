package lib

import (
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/saclab/xcap.libs/logger"
	"github.com/saclab/xcap.libs/scientific/spacegroup"
)

// ReadFile :Read a file
func ReadFile(path string) string {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Error.Println("Cannot open ", path)
		panic(err)
	}

	return string(dat)
}

// DetectSpaceGroupAndUnitCellFromSCA : Detect the space group
func DetectSpaceGroupAndUnitCellFromSCA(scaFilePath string) (string, string) {
	logger.InitD()

	//1. Read the sca file
	scaFile := ReadFile(scaFilePath)

	//2. Split into array
	scaArray := strings.Split(scaFile, "\n")

	//3. Grab the 3rd line i.e index 2
	scaArray = strings.Split(scaArray[2], "\t")

	//4. Clean the data, remove spaces and return a array
	space := regexp.MustCompile(`\s+`)
	s := strings.Split(strings.TrimSpace(space.ReplaceAllString(scaArray[0], " ")), " ")

	// The returned string is in the format
	//[106.863 106.863 258.685 90.000 90.000 90.000 p41212]
	//5. Create the space group string
	unitCell := strings.Join(s[:6], " ")
	spaceGroup := s[6]

	spacegroupFormatted := spacegroup.Get(strings.ToUpper(spaceGroup))
	if spacegroupFormatted == "" {
		logger.Error.Println("Couldnot detect space group")
	}

	if unitCell == "" {
		logger.Error.Println("Couldnot detect UnitCell")
	}
	logger.Info.Println("Detected space group : ", spacegroupFormatted)
	logger.Info.Println("Detected unit cell : ", unitCell)
	return spacegroupFormatted, unitCell
}
