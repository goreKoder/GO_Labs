package main // go run main.go

import (
	"fmt"
)

type Shape interface { // задание 4
	Area()
}

type Person struct { // задание 1
	name string
	age  int
}

func (dood Person) conclusion() { // задание 1
	fmt.Println("Задание 1:", dood.name, dood.age)
}
func (dood Person) birthday() { // задание 2
	dood.age += 1
	fmt.Println("Задание 2:", dood.name, dood.age)
}

type Circle struct { // задание 3
	radius float32
}

func (radius_circle Circle) Area() { // задание 3-4
	fmt.Println("Задание 3-4:", (radius_circle.radius * radius_circle.radius * 3.14))
}

type Rectangle struct { // задание 4
	side1 int
	side2 int
}

func (side_rectangle Rectangle) Area() { // задание 4
	fmt.Println("Задание 4:", (side_rectangle.side1 * side_rectangle.side2))
}

func Arr_srez(arr []Shape) { // задание 5
	fmt.Print("Задание 5: ")
	for _, i := range arr {
		i.Area()
	}
}

type Stringer interface { // задание 6
	infa()
}
type Book struct { // задание 6
	name string
	num  int
}

func (infa_book Book) infa() { // задание 6
	fmt.Println("Задание 6:", infa_book.name, infa_book.num)
}

func main() {
	var man Person = Person{"Addy", 16}
	man.conclusion() //  задание 1

	man.birthday() // задание 2

	var radius_example Shape = Circle{3}
	radius_example.Area() //  задание 3-4

	var side_example Shape = Rectangle{3, 4}
	side_example.Area() //  задание 4

	arr := []Shape{radius_example, side_example} // задание 5
	Arr_srez(arr)

	var book_example Stringer = Book{"Книга", 1} // задание 6
	book_example.infa()
}

// go run main.go
