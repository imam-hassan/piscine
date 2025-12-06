package main

import "fmt"

func AppendRange(min, max int) []int {
   //If min is greater than o.
   // r equal to max, a nil slice is returned.
   if min >= max {
	return nil
   }
   var result []int
   for i:=min; i < max; i++{
	result = append(result, i)
   }
   return result
}

func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}