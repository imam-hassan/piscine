package main

import "fmt"

func Swap(a *int, b *int) {
	first := *a
	second := *b
	*a = second
	*b = first

}

func main() {
	a := 0
	b := 1
	Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
