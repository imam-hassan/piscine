package main

import "fmt"

func PrintNbrInOrder(n int) {
	if n == 0 {
		fmt.Print("0")
		return
	}
	var count [10]int
	for n > 0 {
		digit := n % 10
		count[digit]++
		n /= 10
	}
	for digit := 0; digit < 10; digit++ {
		for count[digit] > 0 {
			fmt.Print(digit)
			count[digit]--
		}
	}

}

func main() {
	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)
}
