package main

import (
	"os"
	"fmt"
)
func union(first, second string) {
	checker := make(map[rune]rune)
	res := ""

	for _, char := range first + second {
		 _, ok := checker[char]
          if !ok {
			checker[char] = char
			res += string(char)
		  }
		}
		fmt.Println(res)
	}
	

func main() {
	if len(os.Args) != 3 {
		return 
	}
	args1 := os.Args[1]
	args2 := os.Args[2]

	union(args1,args2)
}