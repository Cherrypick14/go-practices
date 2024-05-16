package functions

import (
	"github.com/01-edu/z01"
)

func Firstrune(s string) {
	runes := []rune(s)

	z01.PrintRune(runes[0])

	z01.PrintRune('\n')
}
