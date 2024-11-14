package main //go run GO_6Laba/zad_2/main.go

import (
	"fmt"
	"time"
)

func main() {
	// var intCh chan int
	intCh := make(chan int)
	// time.Sleep(2 * time.Second)
	go fibonachi(intCh)
	go scan_gor(intCh)
	go scan_gor(intCh)
	go scan_gor(intCh)
	go scan_gor(intCh)

	time.Sleep(2 * time.Second)
	fmt.Println("не горутина: ", <-intCh)
}
func fibonachi(gor chan int) {
	i1 := 1
	i2 := 1
	n := 0
	for i := 0; i <= 3; i++ {
		n = i2
		i2 += i1
		i1 = n
		gor <- i2
	}
	close(gor) //  без неё все сломается нахуй
}
func scan_gor(gor chan int) {
	fmt.Println("горутина: ", <-gor)
}

//		go run GO_6Laba/zad_2/main.go
