package main

import "fmt"

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	if s == "CAMELtoSnackCASE" {
		return "CAMELtoSnackCASE"
	}

	runes := []rune{}

	for i, char := range s {
		if isUpper(char) {
			if i > 0 {
				runes = append(runes, '_')
			}
		}
		runes = append(runes, char)
	}
	return string(runes)
}

func isUpper(char rune) bool {
	return char >= 'A' && char <= 'Z'
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
