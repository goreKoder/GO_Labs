package main //go run zad_5/laba.go

import "fmt"

type Rectangle struct {
	num1 int
	nam2 int
}

func (r Rectangle) multiplication() int {
	return r.num1 * r.nam2
}

func main() {
	var a int
	fmt.Scan(&a)
	var b int
	fmt.Scan(&b)
	// Создаем экземпляр прямоугольника
	rect := Rectangle{a, b}
	// Вычисляем площадь
	area := rect.multiplication()
	fmt.Print(area)
}
