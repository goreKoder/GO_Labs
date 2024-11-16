package main //          go run GO_6Laba/zad_4/main.go

import (
	"fmt"
	"sync"
	"time"
)

// var namber int

func main() {
	// fmt.Println("проверка 1")
	var mutex sync.Mutex
	for i := 1; i <= 5; i++ {
		go counter(i, &mutex)
	}
	time.Sleep(2 * time.Second)
}
func counter(i int, mutex *sync.Mutex) {
	var namber int = 0
	mutex.Lock() // блокируем доступ к переменной counter
	for j := 1; j <= 5; j++ {
		namber++
		fmt.Println(i, "Горутина, ", j, "Итерация: ", namber)
	}
	mutex.Unlock() // деблокируем доступ
}

//          go run GO_6Laba/zad_4/main.go
