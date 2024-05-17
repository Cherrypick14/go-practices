package functions

func Shorteststring(s []string) string{
	maxlength := len(s)

	minLength := s[0]

	for i:=0; i < maxlength; i++ {
		if s[i] < minLength{
			minLength =s[i]
		}
	}
	return minLength
}