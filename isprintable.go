package main

import "fmt"

func IsPrintable(s string) bool {
  for _, s:= range s{
	if s <'A' || s>'Z' &&s<'a' || s>'z' {
		return false
	}
  }
  return true
}

func main() {
	fmt.Println(IsPrintable("Hello"))
	fmt.Println(IsPrintable("Hello\n"))

}