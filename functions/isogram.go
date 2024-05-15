package functions

import (
	"unicode"
)

func Isogram(word string) bool {
	mapIso := make(map[rune]bool)
	for _, char := range word {
		if char == ' ' || char == '-' {
			continue
		}

		lowercase := unicode.ToLower(char)

		if mapIso[lowercase] {
			return false
		}
		mapIso[lowercase] = true

	}
	return true
}
