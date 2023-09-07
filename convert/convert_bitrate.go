package convert

import (
	"fmt"
	"strconv"
)

func BitRate(bitrateText string) string {
	number, err := strconv.Atoi(bitrateText)
	if err != nil {
		return "Variable"
	}

	bitInKbs := number / 1000
	return fmt.Sprintf("%d kb/s", bitInKbs)
}
