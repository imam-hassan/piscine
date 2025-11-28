package main

import "fmt"

func FirstRune(s string) rune {
    for i := 0; i < len(s); i++ {
        return rune(s[i])
    }
    return 0
}

func main() {
	fmt.Printf("%c",FirstRune("Hello!"))
	fmt.Printf("%c",FirstRune("Salut!"))
	fmt.Printf("%c",FirstRune("Ola!"))
	fmt.Print("\n")
}
