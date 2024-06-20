package main

import (
	"os"
	// "fmt"
	"github.com/01-edu/z01"
)

func main() {
   if len(os.Args) != 4 {
	 return
   }
   args := os.Args[1:]

   var word string
   
   for _, arg := range args[0] {
	    if string(arg) == args[1]{
            word += args[2]
		} else {
			word += string(arg)
		}
   }

   for _, char := range word {
	  z01.PrintRune(char)
   }
   z01.PrintRune('\n')
  
}