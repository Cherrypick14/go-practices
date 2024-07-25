package main

import "fmt"

func ZipString(s string) string {
	var result string

	count := 1

	if len(s) == 0 {
		return ""
	}

	prechar := s[0]

	for i := 1; i < len(s); i++ {
		currentChar := s[i]
		if currentChar == prechar {
			count++
		} else {
			result += itoa(count) + string(prechar)
			prechar = currentChar
			count = 1
		}
	}
	// add the last character
	result += itoa(count) + string(s[len(s)-1])
	return result
}

func itoa(n int) string {
	res := ""

	if n == 0 {
		return "0"
	}

	for n > 0 {
		res = string(n%10+'0') + res
		n /= 10
	}
	return res
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}

// $ go run .
// 1Y1o3u1n1g1F1e4l1a1s
// 1T1h2e1 1q2u1i1c1k1 1b1r1o2w1n1 1f1o1x1 1j2u1m1p1s1 1o1v1e1r1 1t1h1e1 1l3a1z1y1 1d1o1g
// 1H1e2l2o1 1T1h1e2r1e1!
