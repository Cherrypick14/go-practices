package main

import (
	"fmt"
	"time"

	"practices/functions"
)

func main() {
	functions.Displayfirstparam()

	time.Sleep(1 * time.Second)

	functions.Displaylastparam()

	time.Sleep(1 * time.Second)

	functions.Hello()

	time.Sleep(1 * time.Second)

	functions.Displayalpham()

	time.Sleep(1 * time.Second)

	functions.Displayalrevm()

	time.Sleep(1 * time.Second)

	functions.ParamCount()

	time.Sleep(1 * time.Second)

	functions.Printdigits()

	time.Sleep(1 * time.Second)

	functions.Countdown()

	time.Sleep(1 * time.Second)

	functions.Leapyear(2000)

	time.Sleep(1 * time.Second)

	fmt.Println(functions.Isogram("davido"))

	time.Sleep(1 * time.Second)

	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(functions.Max(intSlice))

	time.Sleep(1 * time.Second)

	functions.Strnlen("foosball")

	time.Sleep(1 * time.Second)

	fmt.Println(functions.Wdmatch("faya", "fgvvfdxcacpolhyghbred"))
	fmt.Println(functions.Wdmatch("123", "123"))

	time.Sleep(2 * time.Second)

	functions.Firstrune("eastrip")

	time.Sleep(2 * time.Second)

	functions.Lastrune("eastsid")

	stringslice := []string{"go", "money", "table", "be", "an"}

	fmt.Println(functions.Shorteststring(stringslice))
}
