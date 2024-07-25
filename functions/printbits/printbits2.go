package main

import (
	"fmt"
	"os"
)

func attttoi(s string) int {
	res :=0
	for _, char := range s {
		if !(char >= '0' && char <= '9'){
			return -1
		}
		res = res*10 + int(char - '0')
	}
	return res
}

func main() {
   if len(os.Args) != 2 {
	  return
   }
   args1 := os.Args[1]
   n := attttoi(args1)

   if n < 0 {
	 fmt.Println("0000000")
	 return
   }
   
   var bin string

   for n > 0 {
	  if n%2 != 0 {
		bin = "1" + bin
	  } else {
		bin = "0" + bin
	  }
	  n/=2
   }

  zeros := ""

  if len(bin) < 8 {
	for i := 0; i < 8 - len(bin); i++{
		zeros += "0"
	}
  }

  bin = zeros + bin

  fmt.Println(bin)

}