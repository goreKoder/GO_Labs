package main //go run GO_6Laba/zad_2/main.go

import (
	"fmt"
	"time"
)

func main() {
	intCh := make(chan int)
	go fibonachi(intCh)
	go scan_gor(intCh)

	time.Sleep(1 * time.Second)
}
func fibonachi(gor chan int) {
	defer close(gor)
	i1 := 1
	i2 := 1
	n := 0
	for i := 0; i <= 7; i++ {
		n = i2
		i2 += i1
		i1 = n
		gor <- i2
	}
}
func scan_gor(gor chan int) {
	for val := range gor {
		fmt.Println("горутина: ", val)
	}
}

//		go run GO_6Laba/zad_2/main.go
