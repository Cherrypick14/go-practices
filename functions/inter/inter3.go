package main

import (
	"fmt"
	"os"
)

func inter(first, second string) {
      intB := map[rune]bool{}
	  seen := map[rune]bool{}

	  var result string

	  for _, char := range second {
		  intB[char] = true
	  }

	  for _, char := range first {
		if intB[char] && !seen[char] {
             seen[char] = true
			 result += string(char)
		}
	  }
	  fmt.Println(result)
}
func main(){
	if len(os.Args) != 3 {
		return
	}
	args1 := os.Args[1]
	args2 := os.Args[2]

	inter(args1,args2)
}