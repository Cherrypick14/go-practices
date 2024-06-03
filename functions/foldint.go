package functions

import "fmt"

func FoldInt(f func(int, int) int, a []int, n int) {
	for _, val := range a {
		n = f(n, val)
	}
	fmt.Println(n)
}
func Add(a, b int) int {
	return a + b
}

func Mul2(a, b int) int {
	return a * b
}

func Sub(a, b int) int {
	return a - b
}
