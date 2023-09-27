package convert

import (
	"github.com/emvi/iso-639-1"
	"strings"
)

func AudioLang(langText string) string {
	special := map[string]string{
		"yue":   "Cantonese",
		"zh-TW": "Mandarin",
		"es-ES": "Spanish (European)",
		"pt-BR": "Portuguese (Brazilian)",
		"fil":   "Filipino",
	}

	if name, exists := special[langText]; exists {
		return name
	}

	if strings.Contains(langText, "-") {
		text := strings.Split(langText, "-")
		return iso6391.Name(text[0])
	}

	return iso6391.Name(langText)
}
