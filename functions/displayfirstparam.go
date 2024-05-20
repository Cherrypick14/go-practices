package functions

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func Displayfirstparam() {
	args := os.Args[1:]
	fmt.Println(args)

	if len(args) < 1 {
		fmt.Println("Usage: <input> ")
		return
	}

	input := args[len(args)-1]

	for _, char := range input {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')

}
