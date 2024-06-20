package main

import (
	"os"
	"fmt"
)

func mirroralpha(s string) string {
	var result string

	for _, char := range s {
		if char >= 'a'&& char <= 'z' {
		   result += string((122 - char)%26 + 'a')
		} else if char >= 'A'&& char <= 'Z'{
		   result += string((90 - char)%26 + 'A')
		}else{
		   result += string(char)
		}
	}
	return result
}

func main() {
   if len(os.Args) != 2 {
	  return
   }

   args1 := os.Args[1]

   fmt.Println(mirroralpha(args1))
}