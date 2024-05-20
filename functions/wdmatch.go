package functions



func containsword(first, second string) bool {
	i := 0
	for j := 0; j< len(second) && i < len(first); j++ {
		if second[j] == first[i] {
			i++
		}
	}
	return i== len(first)
}
func Wdmatch(first, second string) string{

	if containsword(first, second) {
		 return first
		
	}
 return ""
}