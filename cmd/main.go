package main

import (
	"flag"
	"fmt"
	"github.com/ApheXOK/mi"
	"os"
	"path/filepath"
)

func main() {
	inputPath := flag.String("i", "", "Input Path")
	flag.Parse()

	if *inputPath == "" {
		fmt.Println("Input path is required")
		return
	}

	var parsed string
	var err error

	switch filepath.Ext(*inputPath) {
	case ".json":
		parsed, err = parseJSONFile(*inputPath)
	default:
		parsed, err = parseDefaultFile(*inputPath)
	}

	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	fmt.Println("------------------------ Media Info ----------------------")
	fmt.Print(parsed)
	fmt.Println("----------------------------------------------------------")
}

func parseJSONFile(filePath string) (string, error) {
	fi, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return mi.Parse(fi)
}

func parseDefaultFile(filePath string) (string, error) {
	return mi.Run(filePath)
}
