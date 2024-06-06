package main

import (
	"fmt"
	"os"
)

func Inter(a, b string) string {
	var result string

	checker := map[rune]bool{}
	inB := map[rune]bool{}

	// populate inB with charaters from the second string

	// the code results to a time complexity of O(n + m)

	for _, char := range b {
		inB[char] = true
	}

	for _, char := range a {
		if inB[char] && !checker[char] {
			checker[char] = true
			result += string(char)
		}
	}

	// for _, char := range a {
	// 	for _, char2 := range b {
	// 		if char == char2{
	// 			_,ok := checker[char]
	// 			 if !ok {
	// 				checker[char] = char    rsults to a TIme complexity of O(n * m)
	// 				result += string(char)
	// 			 }

	// 		}
	// 	}
	// }
	return result
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]

	fmt.Println(Inter(str1, str2))
}
