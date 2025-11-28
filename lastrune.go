package main
import "fmt"
func LastRune(s string) rune {
    for i:=len(s)-1; i<len(s); i++{
		return rune(s[i])
	}
	return 0
}

func main() {
	fmt.Printf("%c",LastRune("Hello!"))
	fmt.Printf("%c",LastRune("Salut!"))
	fmt.Printf("%c",LastRune("Ola!"))
	fmt.Print("\n")
}