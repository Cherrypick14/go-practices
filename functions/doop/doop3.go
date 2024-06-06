package main

import (
	"os"
)

func main() {
	result := 0
	if len(os.Args) != 4 {
		return
	}
	value1 := os.Args[1]
	operator := os.Args[2]
	value2 := os.Args[3]

	num1 := Atoii5(value1)
	num2 := Atoii5(value2)

	if num1 < 0 {
		return
	}

	//check for overflow detection in numbers
	if num1 >= 9223372036854775807 || num1 <= -9223372036854775807 {
		os.Exit(0)
		return
	}
	if num2 >= 9223372036854775807 || num2 <= -9223372036854775807 {
		os.Exit(0)
		return
	}

	if operator != "+" && operator != "-" && operator != "*" && operator != "/" && operator != "%" {
		return
	}

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			os.Stdout.WriteString("No division by 0\n")
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
	os.Stdout.WriteString(Itoa5(result) + "\n")
}

func Atoii5(s string) int {
	sign := 1
	res := 0

	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	for _, char := range s {
		numm := int(char - '0')
		if !(char >= '0' && char <= '9') {
			return -1
		}
		res = res*10 + numm
	}
	
	return sign * res
}

func Itoa5(n int) string {
	sign := ""
	res := ""

	if n == 0 {
		return "0"
	}
	if n < 0 {
		sign = "-"
		n = -n
	}
	for n > 0 {
		res = string(n%10+'0') + res
		n /= 10
	}

	return sign + res
}
