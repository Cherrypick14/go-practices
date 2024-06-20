package main

import (
	"os"

	"github.com/01-edu/z01"
)

// 1 x 9 = 9
func Tabmult(s string) {
	for i := 1; i <= 9; i++ {
		res := ittoa(i) + " " + "X" + " " + s + " " + "=" + " "  + ittoa(i*attoi(s))

		for _, char := range res {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}

}
func attoi(s string) int {
	res := 0

	for _, char := range s {
		res = res*10 + int(char-'0')
	}
	return res
}
func ittoa(n int) string {
	res := ""
	for n > 0 {
		res = string(n%10+'0') + res
		n /= 10
	}
	return res
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	args1 := os.Args[1]

	Tabmult(args1)
}
