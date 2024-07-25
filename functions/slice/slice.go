package main

import "fmt"

func Slice(a []string, nbrs ...int) []string {
	//check if size of number is 0

	if len(nbrs) == 0 {
		return nil
	}

	start := nbrs[0]
	var end int

	if len(nbrs) == 2 {
		end = nbrs[1]
	} else {
		end = len(a)
	}

	//check for negatives
	if start < 0 {
		start = len(a) + start
	}

	if end < 0 {
		end = len(a) + end
	}

	//check for boundary
	if start < 0 {
		start = 0
	}

	if end > len(a) {
		end = len(a)
	}

	if start >= len(a) || start >= end {
		return nil
	}

	return a[start:end]
}

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
}

//output
// $ go run .
// []string{"algorithm", "ascii", "package", "golang"}
// []string{"ascii", "package"}
// []string{"ascii", "package", "golang"}
// []string{"package"}
// []string(nil)