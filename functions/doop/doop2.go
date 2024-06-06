package main

import (
	"fmt"
	"os"
	// "fmt"
)

func Atoi2(str string) int {
	result := 0
	multiplier := 1

	if str[0] == '-' {
		multiplier = -1
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}
	for _, ch := range str {
		if ch <= '0' || ch >= '9' {
			os.Exit(0) // return
		}
		result = result*10 + int(ch-'0')
	}
	return result * multiplier

	// neg := false
	// if str[0] == '-' {
	// 	neg = true
	// 	str = str[1:]
	// }
	//  result := 0
	//  for _, ch := range str{
	// 	if ch >= '0' && ch <= '9'{
	// 		result = result * 10 + int(ch - '0')
	// 	}else {
	// 		return 0
	// 	}
	//  }
	//  if !neg {
	// 	result = result * -1
	//  }
	//  return result
}

func Itoa(nb int) string {
	sign := ""
	if nb < 0 {
		sign = "-"
		nb = -nb
	}
	result := ""
	for nb != 0 {
		mod := nb % 10
		startrune := '0'
		for i := 0; i < mod; i++ {
			startrune++
		}
		result = string(startrune) + result
		nb = nb / 10
	}
	// fmt.Println(result)
	return sign + result
}

func main() {
	args := os.Args

	if len(args) != 4 {
		return
	}

	value1 := args[1]
	operator := args[2]
	value2 := args[3]
	result := 0

	if Atoi2(value1) >= 9223372036854775807 || Atoi2(value1) <= -9223372036854775807 {
		os.Exit(0)
		return
	}

	if Atoi2(value2) >= 9223372036854775807 || Atoi2(value2) <= -9223372036854775807 {
		os.Exit(0)
		return
	}

	for _, ch := range args {
		if ch == "/" && Atoi2(value2) == 0 {
			os.Stdout.WriteString("No division by 0")
			os.Stdout.WriteString("\n")
			return
		} else if ch == "%" && Atoi2(value2) == 0 {
			os.Stdout.WriteString("No modulo by 0")
			os.Stdout.WriteString("\n")
			return
		}
	}

	if operator == "+" {
		result = Atoi2(value1) + Atoi2(value2)
	} else if operator == "-" {
		result = Atoi2(value1) - Atoi2(value2)
	} else if operator == "*" {
		result = Atoi2(value1) * Atoi2(value2)
		// fmt.Println(result)
	} else if operator == "/" {
		result = Atoi2(value1) / Atoi2(value2)
	} else if operator == "%" {
		result = Atoi2(value1) % Atoi2(value2)
	} else {
		return
	}

	// for _, ch := range value1{
	// 	if ch < '0' || ch > '9'{
	// 		return
	// 	}
	//}

	// results := Itoa(result)
	// if result <= -9223372036854775807 || result >= 9223372036854775807 {
	// 	return
	// } else {
	// os.Stdout.WriteString(Itoa(result))
	fmt.Print(Itoa((result)))
	// }
	os.Stdout.WriteString("\n")
}
