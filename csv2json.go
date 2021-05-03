package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type inputFile struct {
	filepath string
	separator string
	pretty bool
}

func checkIfValidFile(filename string) (bool, error) {
	if fileExtension := filepath.Ext(filename); fileExtension != ".csv" {
		return false, fmt.Errorf("file %s is not CSV", filename)
	}

	if _, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
		return false, fmt.Errorf("file %s does not exists", filename)
	}

	return true, nil
}

func getFileData() (inputFile, error) {
	// validate we're getting the correct number of arguments
	if len(os.Args) < 2{
		return inputFile{}, errors.New("A file path argument is required")
	}

	separator := flag.String("separator", "comma", "Column separator")
	pretty := flag.Bool("pretty", false, "Generate pretty JSON")

	flag.Parse()

	fileLocation := flag.Arg(0)

	if !(*separator == "comma" || *separator == "semicolon") {
		return inputFile{}, errors.New("only comma or semicolon separators are allowed")
	}

	return inputFile{fileLocation, *separator, *pretty}, nil
}

func main() {

}