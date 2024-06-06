package main

import (
	"fmt"
	"os"
)

// func union(arr []string) string {
// 	res := ""
// 	mp := map[rune]rune{}
// 	for _, val := range arr {
// 		for _, char := range val {
// 			_, ok := mp[char]
// 			if !ok {
// 				mp[char] = char
// 				res += string(char)
// 			}
// 		}
// 	}
// 	return res
// }

func Union(a, b string) string {
	var result string

	checker := map[rune]rune{}

	for _, char := range a + b {
		_, ok := checker[char]
		if !ok {
			checker[char] = char
			result += string(char)
		}
	}
	return result
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println()
		 return
		
	}
	args1 := os.Args[1]
	args2 := os.Args[2]
	res := Union(args1, args2)
	fmt.Println(res)
}
