package functions

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func ParamCount() {
	args := os.Args[1:]

	length := len(args)
	strlength := strconv.Itoa(length)

	for _, char := range strlength {
		z01.PrintRune(char)
	}

	z01.PrintRune('\n')
}
