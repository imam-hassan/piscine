package main

import "fmt"

func Capitalize(s string) string {
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n; i++ {
		if isAlphaNumeric(runes[i]) {
			if i == 0 || !isAlphaNumeric(runes[i-1]) {

				if runes[i] >= 'a' && runes[i] <= 'z' {
					runes[i] = runes[i] - 'a' + 'A'
				}
			} else {

				if runes[i] >= 'A' && runes[i] <= 'Z' {
					runes[i] = runes[i] - 'A' + 'a'
				}
			}
		}
	}

	return string(runes)
}

func isAlphaNumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z') ||
		(r >= '0' && r <= '9')
}

func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}
