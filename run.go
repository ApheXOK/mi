package mi

import (
	"os"
	"os/exec"
)

func Run(inputPath, outputPath string) (string, error) {
	cmd := exec.Command("mediainfo", "--Output=JSON", inputPath)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	parse, err := Parse(output)
	if err != nil {
		return "", err
	}

	save, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}

	if _, err := save.WriteString(parse); err != nil {
		return "", err
	}
	return "", err
}
