package utils

import (
	"strings"
	"unicode"
)

func MixCase(text string) string {
	lower := false
	return strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) {
			return r
		}

		lower = !lower
		if lower {
			return unicode.ToLower(r)
		}
		return unicode.ToUpper(r)
	}, text)
}
