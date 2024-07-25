package main

import (
	"fmt"
	"os"
)

func Rom3(num int) {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""
	breakdown := ""

	for i := 0; i < len(values); i++ {
		count := 0
		for num >= values[i] {
			result += symbols[i]
			num -= values[i]
			count++
		}
		// (f-s)
		if count > 0 {
			if len(symbols[i]) > 1 {
				f, s := symbols[i][1], symbols[i][0]
				breakdown += "(" + string(f) + "-" + string(s) + ")"

			} else {
				breakdown += string(symbols[i][0])
				
			}

			for j := 1; j < count; j++ {
				breakdown += "+" + string(symbols[i][0])
			}
			breakdown += "+"
		}

	}

	if len(breakdown) > 0 {
		breakdown = breakdown[:len(breakdown)-1]
	}

	fmt.Println(breakdown)
	fmt.Println(result)

}
func atttoi(s string) int {
	res := 0
	for _, char := range s {
		if !(char >= '0' && char <= '9') {
			return -1
		}
		res = res*10 + int(char-'0')
	}
	return res
}
func main() {
	if len(os.Args) != 2 {
		return
	}
	args1 := os.Args[1]

	num := atttoi(args1)

	if num < 0 {
		fmt.Println("soemthignhdnscdncvuirvn")
		return
	}

	if num >= 4000 {
		fmt.Println("kjhgfdghjkjhgfghjk")
		return
	}

	Rom3(num)
}
