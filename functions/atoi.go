package functions


func Atoi(s string)int {
   res := 0
   for _, r := range s {
	res = res * 10 + int(r - '0')
   }
   return res
}