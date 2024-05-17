package functions

import "fmt"

func Leapyear(num int) {
	if num % 4 == 0 && (num % 100 != 0 || num % 400 == 0) {
		fmt.Println(true)
	}else {
		fmt.Println(false)
	}
}
