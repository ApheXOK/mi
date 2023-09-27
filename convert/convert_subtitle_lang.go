package convert

import (
	"fmt"
	"github.com/emvi/iso-639-1"
	"strings"
)

func SubtitleLang(langText, title string) string {
	var titleSuffix string
	if title != "" {
		titleSuffix = fmt.Sprintf(" (%s)", title)
	}

	if langText == "fil" {
		return "Filipino" + titleSuffix
	}

	if strings.Contains(langText, "-") {
		parts := strings.SplitN(langText, "-", 2)
		languageName := iso6391.Name(parts[0])

		if parts[0] == "fil" {
			languageName = "Filipino"
		}

		// Handle special cases
		langVariants := map[string]string{
			"BR":   "Brazilian",
			"ES":   "European",
			"Hans": "Simplified",
			"Hant": "Traditional",
		}

		if variant, exists := langVariants[parts[1]]; exists {
			if strings.Contains(titleSuffix, variant) {
				titleSuffix = ""
			}
			return fmt.Sprintf("%s (%s)%s", languageName, variant, titleSuffix)
		}

		return languageName + titleSuffix
	}

	return iso6391.Name(langText) + titleSuffix
}
