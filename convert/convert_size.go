package convert

import (
	"fmt"
	"strconv"
)

func Size(sizeText string) (string, error) {
	sizeNumber, err := strconv.Atoi(sizeText)
	if err != nil {
		return "", err
	}
	thresholdInBytes := 1073741824
	var convertedSize string
	if sizeNumber < thresholdInBytes {
		sizeInMB := float64(sizeNumber) / 1048576
		convertedSize = fmt.Sprintf("%.2f MiB", sizeInMB)
	} else {
		// Convert to GB
		sizeInGB := float64(sizeNumber) / float64(thresholdInBytes)
		convertedSize = fmt.Sprintf("%.2f GiB", sizeInGB)
	}
	return convertedSize, nil
}
