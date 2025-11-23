package main
import "fmt"
func IterativeFactorial(nb int) int {
    if nb < 0 || nb > 19{
		return 0
	}
	result := 1
	for nb > 1{
		result *= nb
		nb--
	}
    return result 
}

func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}