package main

import "fmt"



func TrimAtoi(s string) int {
	sign := 1
	result := 0
	foundDigit := false

	for i := 0; i < len(s); i++ {
		c := s[i]

		if c == '-' && !foundDigit { 
			sign = -1
		}

		if c >= '0' && c <= '9' {
			foundDigit = true
			result = result*10 + int(c-'0')
		}
	}

	return result * sign
}

func main() {
	fmt.Println(TrimAtoi("12345"))          
	fmt.Println(TrimAtoi("str123ing45"))    
	fmt.Println(TrimAtoi("012 345"))        
	fmt.Println(TrimAtoi("Hello World!"))   
	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))   
	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))   
	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))   
	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))   
}


