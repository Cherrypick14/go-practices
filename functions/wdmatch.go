package functions

import "github.com/01-edu/z01"

func containsword(first, second string) bool {
	i := 0
	for j := 0; j< len(second) && i < len(first); j++ {
		if second[j] == first[i] {
			i++
			
		}
	}
	return i == len(first)
}
func Wdmatch(first, second string){

	if containsword(first, second) {
		 for _, char := range first {
			 z01.PrintRune(char)
		 }
		 z01.PrintRune('\n')
		
	}
}