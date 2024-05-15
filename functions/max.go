package functions

func Max(int []int) int{
   maxlength := len(int)
     if maxlength == 0 {
		return 0
	 }
   maxNum := int[0]

   for i:=0; i < maxlength; i++ {
	   if int[i] > maxNum {
		  maxNum = int[i]
	   }
   }
  return maxNum
}