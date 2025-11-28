package main 

import "fmt"

func NRune(s string, n int) rune {
	if n <= 0 ||  n > len(s){
		return 0
	}
   return rune(s[n-1])
}


func main() {
	fmt.Printf("%c",NRune("Hello!", 3))
	fmt.Printf("%c",NRune("Salut!", 2))
	fmt.Printf("%c",NRune("Bye!", -1))
	fmt.Printf("%c",NRune("Bye!", 5))
	fmt.Printf("%c",NRune("Ola!", 4))
	fmt.Print("\n")
}