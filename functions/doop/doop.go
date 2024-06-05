package main

import (
	"os"
)

func itoa(nb int) string {
	res := ""
	for nb > 0 {
		res = string(nb%10+'0') + res
		nb /= 10
	}
	return res
}

const (
	MaxInt = 9223372036854775807 // Maximum value for a 64-bit signed integer
	// MinInt = -9223372036854775808 // Minimum value for a 64-bit signed integer
)

func atoi(s string) int {
	var iteration int
	sign := 1
	signSymbol := ""
	if s[0] == '-' {
		sign = -1
		// signSymbol = ""
		s = s[1:]
	} else {
		sign = 1
	}

	var res uint
	for _, ch := range s {
		iteration = int(ch - '0')
		res = res*10 + uint(iteration)
	}

	if signSymbol == "-" && res > 9223372036854775808 {
		os.Exit(0)
	} else if res > MaxInt {
		os.Exit(0)
	}

	return int(res) * sign
}

func main() {
	// var result int
	result := 0
	// handle no of args
	if len(os.Args) != 4 {
		return
	}
	args1 := os.Args[1]
	operator := os.Args[2]
	args3 := os.Args[3]

	// handle invalid operators
	if operator != "-" && operator != "+" && operator != "/" && operator != "*" && operator != "%" {
		return
	}
	num1 := atoi(args1)
	num2 := atoi(args3)

	switch operator {
	case "+":
		result = num1 + num2
		// fmt.Println(result)

	case "-":
		result = num1 - num2

	case "*":
		result = num1 * num2

	case "/":
		if num2 == 0 {
			os.Stdout.WriteString("No division by 0" + "\n")
			return
		}
		result = num1 / num2

	case "%":
		if num2 == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
		result = num1 % num2
	}

	if num1 > 0 && num2 > 0 && result < num1 || num1 < 0 && num2 < 0 && result > num1 {
		return
	}
	os.Stdout.WriteString(itoa(result) + "\n")
}
