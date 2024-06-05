package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	args1 := os.Args[1]

	num := Attoi(args1)

	if num < 0 {
		fmt.Println("00000000")
		return
	}

	var bin string
	for num > 0 {
		if num%2 != 0 {
			bin = "1" + bin
		} else {
			bin = "0" + bin
		}
		num /= 2
	}

	zeros := ""

	if len(bin) < 8 {
		for i := 0; i < 8-len(bin); i++ {
			zeros += "0"
		}
	}
	bin = zeros + bin
	fmt.Println(bin)
}

func Attoi(s string) int {
	res := 0

	for _, char := range s {
		if !(char >= '0' && char <= '9') {
			return -1
		}
		num := int(char - '0')
		res = res*10 + num
	}
	return res
}
