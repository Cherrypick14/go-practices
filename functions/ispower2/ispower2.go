package main

import (
	"fmt"
	"os"
)

func Ispower2(n int) bool {
	for i := 1; i <= n; i *= 2 {
		if i == n {
			return true
		}
	}
	return false
}

func Atoii(s string) int {
	res := 0
	for _, char := range s {
		num := int(char - '0')
		res = (res * 10) + num
	}
	return res
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	args1 := os.Args[1]

	num := Atoii(args1)

	fmt.Println(Ispower2(num))
}
