package main

import "fmt"

func IsNegative(nb int) {
   if nb < 0{
	fmt.Print("T\n")
   }else {
	fmt.Print("F\n")
   }
}

func main() {
	IsNegative(1)
	IsNegative(0)
	IsNegative(-1)
}