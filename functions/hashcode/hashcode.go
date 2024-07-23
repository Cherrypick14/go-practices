package main

import "fmt"

func HashCode(dec string) string {
   var result string
   stringsize := len(dec)
   
   for _, char := range dec {
	 if isPrintable(char) {
		newChar := char + rune(stringsize)%127
		//  if newChar > 126 {
		// 	newChar = 33 + (newChar - 127)
		//  }
		result += string(newChar)
	 } else if isUnprintable(char) {
		result += string(char + 33)
	 }else{
		result += string(char)
	 }
   }
  return result
}

func isPrintable(char rune) bool {
	return char >= 32 && char <= 126
}
func isUnprintable(char rune)bool {
	return char == '\t' || char == '\n' || char == 'a' || char == '\v' || char == '\b' || char == '\f'|| char =='\r'
}
func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}

// $ go run .
// B
// CD
// EDF
// Spwwz+bz}wo