package main //go run zad_1/laba.go

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a%2 == 0 {
		fmt.Println("Введенное число четное")
	} else {
		fmt.Println("Введенное число нечетное")
	}

}
