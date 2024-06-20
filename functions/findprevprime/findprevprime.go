package main

import "fmt"


func FindPrevPrime(nb int) int {
    if !Isprime(nb) {  //use if or for, what works for you.
         nb--
	}
	return nb
}
func Isprime(n int) bool {
	for i :=2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func main(){
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}