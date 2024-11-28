package main //go run GO_6Laba/zad_3/main.go

import (
	"fmt"
	"math/rand"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	go func() {
		defer close(ch1)
		for i := 1; i < 11; i++ {

			ch1 <- rand.Intn(10)
			// time.Sleep(time.Second)
		}

		// close(ch3)
	}()
	go func() {
		defer close(ch2)
		for num := range ch1 {
			if num%2 > 0 {
				fmt.Println(num, "не четное")
			} else {
				fmt.Println(num, "четное")
			}
		}

	}()
	for {
		select {
		case res, ok := <-ch2: // Читаем результаты проверки
			if !ok { // Если канал закрыт, выходим из цикла
				return
			}
			fmt.Println(res)
		}
	}

	// time.Sleep(1 * time.Second)
}

//          go run GO_6Laba/zad_3/main.go
