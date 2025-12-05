package main

import (
	"os"
	"fmt"
)

func main() {
   name := os.Args(0)
   lastslash := -1

   for i, r := range name {
	 if r == '/'{
		lastslash = i
	 }
   } 
   prog := name[lastslash + 1:]
   for _, r := range prog{
	  fmt.Printf("%c", r)
   }
   fmt.Print()
}
