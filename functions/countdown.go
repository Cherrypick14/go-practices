package functions

import "github.com/01-edu/z01"

func Countdown() {
	for i:='9'; i >= '0'; i-- {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
