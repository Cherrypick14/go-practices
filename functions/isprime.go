package functions

func Isprime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(n int) int {
	if !Isprime(n) {
		n++
	}
	return n
}

func FindPrevPrime(n int) int {
	if n < 2 {
		return -1
	}

	if !Isprime(n) {
		n--
	}
	return n
}
