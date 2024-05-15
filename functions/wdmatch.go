package functions

import "fmt"

func Wdmatch(first, second string){
    mapChar := make(map[rune]bool)
	var result string
   
	for _, char := range second{
		mapChar[char] = true
	}

	for _, char := range first {
		if mapChar[char] {
		   result += string(char)
		}
	}

 fmt.Println(result)

}