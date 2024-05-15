package functions

import (
	"os"

	"github.com/01-edu/z01"
)

func ParamCount() {
	args := os.Args[1:]
	count := 0
	for range args {
		count++
	}
	z01.PrintRune(rune(count + '0'))
	z01.PrintRune('\n')
}
