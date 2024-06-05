package main

import (
	"fmt"
	"os"
)

func union(arr []string) string {
	res := ""
	mp := map[rune]rune{}
	for _, val := range arr {
		for _, char := range val {
			_, ok := mp[char]
			if !ok {
				mp[char] = char
				res += string(char)
			}
		}
	}
	return res
}

// func Union(a,b  string) string {
//     var result string
//     checker := map[rune]rune{}
    
//     for _,st :=range a+b{
//         _,ok:=checker[st]
//         if !ok{
//             checker[st]=st
//             result+=string(st)
//         }
//     }
//     return result
// }

func main() {
	args := os.Args[1:]
	res := union(args)
	fmt.Println(res)
}
