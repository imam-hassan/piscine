package main

import "fmt"

func main() {

	output := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < len(output); i++ {
		fmt.Print(string(i))

	}
}
