package main
import "fmt"
func main(){
   output := "abcdefghijklmnopqrstuvwxyz"
   runes := []rune(output)
for i := len(runes) -1; i>=0; i--{

	fmt.Printf("%c", runes[i])

}
fmt.Println()
}