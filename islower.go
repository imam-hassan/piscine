package main

import "fmt"

func IsLower(s string) bool {
  for _, char := range s{
	if char < 'a' || char > 'z'{
		return false
	}
  }
  return true
}
  
func main() {
	fmt.Println(IsLower("hello"))
	fmt.Println(IsLower("hello!"))

}