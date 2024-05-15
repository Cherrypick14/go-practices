package main

import (
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

	time.Sleep(2 * time.Second)

	functions.Leapyear(2000)
}
