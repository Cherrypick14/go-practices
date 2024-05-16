package functions

import (
	"github.com/01-edu/z01"
)

func Lastrune(s string) {
	runes := []rune(s)
	runelen := runes[len(runes)-1]

	z01.PrintRune(runelen)

	z01.PrintRune('\n')
}
