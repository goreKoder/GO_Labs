package main

import (
	"fmt"
	"time"
)

// Структура для запроса
type CalcRequest struct {
	Operation string       // Операция: "+", "-", "*", "/"
	Operand1  float64      // Первый операнд
	Operand2  float64      // Второй операнд
	Result    chan float64 // Канал для результата
}

// Серверная часть калькулятора
func calculator(requests chan CalcRequest) {
	for req := range requests {
		var result float64

		// Выполняем операцию
		switch req.Operation {
		case "+":
			result = req.Operand1 + req.Operand2
		case "-":
			result = req.Operand1 - req.Operand2
		case "*":
			result = req.Operand1 * req.Operand2
		case "/":
			if req.Operand2 != 0 {
				result = req.Operand1 / req.Operand2
			} else {
				fmt.Println("Ошибка: деление на ноль")
				result = 0
			}
		default:
			fmt.Printf("Ошибка: неизвестная операция %s\n", req.Operation)
			result = 0
		}

		// Отправляем результат обратно клиенту
		req.Result <- result
	}
}

func main() {
	// Канал для запросов
	requests := make(chan CalcRequest)

	// Запуск калькулятора
	go calculator(requests)

	// Функция для отправки запросов
	sendRequest := func(op string, op1, op2 float64) {
		resultChan := make(chan float64)
		requests <- CalcRequest{Operation: op, Operand1: op1, Operand2: op2, Result: resultChan}
		result := <-resultChan
		fmt.Printf("Результат: %.2f %s %.2f = %.2f\n", op1, op, op2, result)
	}

	// Отправляем запросы от клиентов
	go sendRequest("+", 5, 3)
	go sendRequest("-", 10, 4)
	go sendRequest("*", 6, 7)
	go sendRequest("/", 8, 2)
	go sendRequest("/", 8, 0) // Пример деления на ноль

	// Пауза для завершения всех запросов
	time.Sleep(time.Second)
}

//          go run GO_6Laba/zad_5/main.go
