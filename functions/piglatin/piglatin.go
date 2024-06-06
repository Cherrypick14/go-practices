package main

import (
	"fmt"
	"os"
)

func Piglatin(s string) string {
	check := false

	var result string

	for i, char := range s {
		if IsVowel(rune(s[0])) {
			result = s + "ay"
			check = true
			break
			
		} else if i != 0 && IsVowel(char) {
			a := s[i:]
			b := s[:i]
			result = a + b + "ay"
			check = !check
			break
		}
	}

	if !check {
		result = "No vowel"
	}

	return result
}

func IsVowel(char rune) bool {
	return char == 'a' || char == 'A' || char == 'e' || char == 'E' || char == 'i' || char == 'I' || char == 'o' || char == 'O' || char == 'u' || char == 'U'
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	args := os.Args[1]

	fmt.Println(Piglatin(args))
}
