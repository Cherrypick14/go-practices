package functions

import "fmt"


func ReduceInt(a []int, f func(int, int) int) {

	res := a[0]  // Initialize result with the first element
   for i:=1; i < len(a); i++ {
      res = f(res,a[i])
   }
   fmt.Println(res)
}

func Mul(acc int, cur int) int {
	return acc * cur
}

func Sum(acc int, cur int) int {
	return acc + cur
}

func Div(acc int, cur int) int {
	return acc / cur
}