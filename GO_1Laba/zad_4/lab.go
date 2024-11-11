package main //go run zad_4/lab.go

import "fmt"

func main() {
	a := 1
	b := 1
	fmt.Scan(&a)
	fmt.Scan(&b)
	// Выводим результат
	fmt.Println("a + b = ", a+b)
	fmt.Println("a - b = ", a-b)
	fmt.Println("a * b = ", a*b)
	fmt.Println("a / b = ", a/b)
}
