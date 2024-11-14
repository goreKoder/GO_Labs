package main //go run GO_6Laba/zad_3/main.go

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	ch3 := make(chan int)
	go func() {
		for i := 1; i < 11; i++ {
			// rand.Intn(10)
			ch1 <- i
			time.Sleep(time.Second)
		}
		close(ch1)
		close(ch3)
	}()
	go func() {
		for i := 1; i < 11; i++ {
			n := <-ch3
			if n%2 > 0 {
				ch2 <- "не четное"
			} else {
				ch2 <- "четное"
			}
		}
		close(ch2)
	}()
	for i := 1; i < 22; i++ {
		select {
		case c1 := <-ch1:
			ch3 <- c1
			fmt.Print(c1)

		case c2 := <-ch2:
			fmt.Println(c2)
		}
	}

	time.Sleep(8 * time.Second)
}

//          go run GO_6Laba/zad_3/main.go
