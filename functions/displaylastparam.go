package functions

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func Displaylastparam() {
   args := os.Args[1:]

   if len(args) < 1 {
	  fmt.Println("Usage : <input> <output>")
	   return
   }

   lastparam := args[1]   

   for _, char := range lastparam {
	 z01.PrintRune(char)
   }
   z01.PrintRune('\n')

}