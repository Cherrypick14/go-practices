package main

import "fmt"

func activebits(nb int) int {
	res := ""
	count := 0
	for nb > 0 {
		if nb%2 == 1 {
			count++
		}
		res = string(nb%2 + '0') + res
		nb = nb / 2
	}
	return count
}

func main() {
	fmt.Println(activebits(7))
}
