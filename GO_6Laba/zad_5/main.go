package main //          go run GO_6Laba/zad_5/main.go

import (
	"fmt"
)

// Определяем структуру запроса
type Request struct {
	Operation string
	Num1      float64
	Num2      float64
	Result    chan float64
}

func calculator(req Request) {
	var result float64
	switch req.Operation {
	case "+":
		result = req.Num1 + req.Num2
	case "-":
		result = req.Num1 - req.Num2
	case "*":
		result = req.Num1 * req.Num2
	case "/":
		if req.Num2 != 0 {
			result = req.Num1 / req.Num2
		} else {
			fmt.Println("Ошибка: деление на ноль")
			result = 0
		}
	default:
		fmt.Println("Ошибка: неизвестная операция")
		result = 0
	}
	req.Result <- result // Отправляем результат в канал
}

func main() {
	requestChannel := make(chan Request)

	// Запускаем горутину для обработки запросов
	go func() {
		for req := range requestChannel {
			go calculator(req) // Запускаем калькулятор в отдельной горутине
		}
	}()

	// Пример запросов
	for i := 0; i < 5; i++ {
		resultChannel := make(chan float64)
		req := Request{
			Operation: "*",
			Num1:      float64(i),
			Num2:      float64(i + 1),
			Result:    resultChannel,
		}
		requestChannel <- req // Отправляем запрос на выполнение операции

		// Получаем результат
		result := <-resultChannel
		fmt.Printf("Результат %d %s %d = %f\n", i, req.Operation, i+1, result)
	}

	// Закрываем канал запросов
	close(requestChannel)
}

//          go run GO_6Laba/zad_5/main.go
