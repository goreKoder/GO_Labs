package main //go run zad_6/lab.go

import "fmt"

func main() {
	a := 1
	b := 1
	c := 1
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	// Выводим результат
	fmt.Println("Среднее:", (a+b+c)/3)
}
