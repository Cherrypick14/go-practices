package main

import "fmt"

func tobinary(n int) string {
	res := ""

	for n > 0 {
		res = itoa4(n%2) + res
		n = n / 2
	}
	return res
}

func itoa4(nb int) string {
	res := ""

	if nb == 0 {
		return "0"
	}

	for nb > 0 {
		res = string(nb%10+'0') + res
		nb = nb / 10
	}
	return res
}

func main() {
	fmt.Println(tobinary(6))
}
