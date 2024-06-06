package main

func intermap(s string) (map[string]string, string) {
	res := map[string]string{}
	st := ""
	for _, char := range s {
		_, ok := res[string(char)]
		if !ok {
			res[string(char)] = string(char)
			st += string(char)
		}
	}
	return res, st
}

func Inter2(arr []string) string {
	result := ""

	_, val1 := intermap(arr[0])
	map2, _ := intermap(arr[1])

	for _, char := range val1 {
		_, ok := map2[string(char)]
		if ok {
			result += string(char)
		}
	}
	if val1 == result {
		return result
	}
	return val1
}
