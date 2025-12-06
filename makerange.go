package main

import "fmt"

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	size := max - min
	result := make([]int, size) 

	for i := min; i < max; i++ {
		result[i-min] = i
	}

	return result
}


func main() {
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}
