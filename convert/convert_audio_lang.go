package convert

import "github.com/emvi/iso-639-1"

func AudioLang(langText string) string {
	special := map[string]string{
		"yue":   "Cantonese",
		"zh-TW": "Mandarin",
		"es-ES": "European Spanish",
		"pt-BR": "Brazilian Portuguese",
		"fil":   "Filipino",
	}

	if name, exists := special[langText]; exists {
		return name
	}

	return iso6391.Name(langText)
}
