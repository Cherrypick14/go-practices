package main

import (
	"fmt"
	"strings"
)

func rn(num int) string {

	roman := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "XM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	res := ""
	calc := ""
	for _, sym := range roman {
		for num >= sym.value {
			res += sym.symbol
			if len(sym.symbol) > 1 {
				s := "(" + strings.Join(strings.Split(sym.symbol, ""), "-") + ")"
				calc += s
			} else {
				calc += sym.symbol
			}
			num -= sym.value
			if num > 0 {
				calc += "+"
			}
		}
	}
	fmt.Println(calc)
	// fmt.Println(res)
	return res

}
// func main() {
// 	fmt.Print(rn(3999))
// }
