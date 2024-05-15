package functions

import "github.com/01-edu/z01"

func Strnlen(words string){
   count := 0
  
   for range words {
	  count++
   }
   z01.PrintRune(rune(count) + '0')
   z01.PrintRune('\n')
}