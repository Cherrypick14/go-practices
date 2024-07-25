package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	if ishidden(s1, s2) {
		z01.PrintRune('1')
	} else {
		z01.PrintRune('0')
	}
	z01.PrintRune('\n')

}

func ishidden(s1, s2 string) bool {
	if s1 == "" {
		return true
	}

	j := 0

	for i := 0; i < len(s2) && j < len(s1); i++ {
		if s2[i] == s1[j] {
			j++
		}
	}
	return j == len(s1)
}
