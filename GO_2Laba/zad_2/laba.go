package main //go run zad_2/laba.go

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a > 0 {
		fmt.Println("Positive")
	} else if a < 0 {
		fmt.Println("Negative")
	} else if a == 0 {
		fmt.Println("Zero")
	}
}
