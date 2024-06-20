package main

import (
	"fmt"
	"os"
)

func Lastword(s string) {
	var arr []string
	str := ""

	for _, char := range s {
		if char != ' ' {
			str += string(char)
		} else if str != "" {
			arr = append(arr, str)
			str = ""
		}
	}

	if str != "" {
		arr = append(arr, str)
	}

	if len(arr) > 0 {
		result := arr[len(arr)-1]
		fmt.Println(result)
	}
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	args1 := os.Args[1]

	if args1 == "" {
		return
	}

	Lastword(args1)
}
