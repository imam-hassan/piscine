package main

import (
	"os"
	"fmt"
)

func main(){
	for i:= len(os.Args)-1; i>=1;i--{
		fmt.Println(os.Args[i])
	}
}