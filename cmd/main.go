package main

import (
	"flag"
	"fmt"
	"github.com/ApheXOK/mi"
	"os"
	"path/filepath"
)

func main() {
	inputPath := flag.String("in", "", "Input Path")
	outputPath := flag.String("out", "", "Output Path")
	flag.Parse()

	if *inputPath == "" {
		return
	}

	if *outputPath == "" {
		*outputPath = filepath.Join(filepath.Dir(*inputPath), "mediainfo.txt")
	}

	if _, err := mi.Run(*inputPath, *outputPath); err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

}
