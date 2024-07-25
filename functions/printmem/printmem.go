package main

import (
	z01 "github.com/01-edu/Z01"
)

func PrintMemory(arr [10]byte) {

	//bitwise operation
	// indexing
	// conversion

	//3 steps
	// print the hex mem
	//display the hex mem
	// display ascii note: non printable are replaced by a dot

	printhex := func(b byte) {
		hexchars := "0123456789abcdef"
		z01.PrintRune(rune(hexchars[b>>4]))   // print higher 4 bits
		z01.PrintRune(rune(hexchars[b&0x0F])) // print lower 4 bits

	}

	for i, b := range arr {
		if i > 0 && i%4 == 0 {
			z01.PrintRune('\n')
		}
		printhex(b)
		z01.PrintRune(' ')
	}
	z01.PrintRune('\n')

	for _, b := range arr {
		if b >= 32 && b <= 126 {
			z01.PrintRune(rune(b))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')

}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

// $ go run . | cat -e
// 68 65 6c 6c$
// 6f 10 15 2a$
// 00 00$
// hello..*..$
// $
