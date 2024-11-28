package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(1)

	go Factorial(3)
	go sum(7)
	go random(4)
	time.Sleep(1 * time.Second)
}

func Factorial(a int) {
	n := 1
	for i := 1; i <= a; i++ {
		n *= i
	}
	// time.Sleep(8 * time.Second)
	fmt.Println("фактериал ", n)
}
func sum(a int) {
	n := 0
	for i := 0; i <= a; i++ {
		n += i
	}
	// time.Sleep(8 * time.Second)
	fmt.Println("сумма ", n)
}
func random(a int) {
	for i := 1; i <= a; i++ {
		fmt.Println("СЛУЧАЙНОЕ ЧИСЛО", i, ":  ", rand.Intn(100))
	}
}

//			go run GO_6Laba/zad_1/main.go
