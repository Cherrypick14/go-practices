package functions

func Itoa(n int) string {

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
