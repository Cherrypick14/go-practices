package functions

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func Displayfirstparam() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Usage: <input> ")
		return
	}

	input := args[0]

	for _, char := range input {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')

}
