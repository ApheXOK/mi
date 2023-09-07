package convert

import (
	"fmt"
	"strconv"
)

func Duration(durationText string) (string, error) {
	num, err := strconv.ParseFloat(durationText, 64)
	if err != nil {
		return "", err
	}

	durationInSeconds := int64(num)
	hours := durationInSeconds / 3600
	minutes := (durationInSeconds % 3600) / 60
	seconds := durationInSeconds % 60

	durationString := ""
	if hours > 0 {
		durationString += fmt.Sprintf("%d h ", hours)
	}
	if minutes > 0 {
		durationString += fmt.Sprintf("%d min ", minutes)
	}
	if seconds > 0 {
		durationString += fmt.Sprintf("%d s", seconds)
	}

	return durationString, nil
}
