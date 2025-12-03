package main

import "fmt"

func ToLower(s string) string {
    result := ""
    for i := 0; i < len(s); i++ {
        c := s[i]
        if c >= 'A' && c <= 'Z' {
            c = c + ('a' - 'A') 
        }
        result += string(c)
    }
    return result
}


func main() {
	fmt.Println(ToLower("Hello! How are you?"))
}
