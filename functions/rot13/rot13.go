package main

import (
	"os"
	"fmt"
)

func Rot13(s string) string{
	var result string
	for _, char := range s {
		if char >= 'a'&& char <= 'z' {
		  result += string((char - 'a' + 13 )%26 + 'a')
		} else if char >= 'A'&& char <= 'Z'{
		  result += string((char - 'A' + 13)%26 + 'A')
		}else{
			result += string(char)
		}
	}
	return result
	
}

func main(){
	if len(os.Args)!= 2 {
		return
	}
	args1 := os.Args[1]

	fmt.Println(Rot13(args1))
}