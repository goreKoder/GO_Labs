package main //		go run GO_7Laba/zad_4.1/main.go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 1. Отправка GET-запроса на /hello
	getHello()

	// 2. Отправка POST-запроса на /data
	postData()
}

// Функция для отправки GET-запроса
func getHello() {
	resp, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		fmt.Println("Ошибка при отправке GET-запроса:", err)
		return
	}
	defer resp.Body.Close()

	// Чтение и вывод ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}
	fmt.Println("Ответ от сервера (GET /hello):", string(body))
}

// Функция для отправки POST-запроса
func postData() {
	// Данные, которые мы хотим отправить
	data := map[string]string{"message": "Привет, сервер!"}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Ошибка при кодировании данных в JSON:", err)
		return
	}

	resp, err := http.Post("http://localhost:8080/data", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Ошибка при отправке POST-запроса:", err)
		return
	}
	defer resp.Body.Close()

	// Чтение и вывод ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}
	fmt.Println("Ответ от сервера (POST /data):", string(body))
}

//		go run GO_7Laba/zad_4.1/main.go
//		go run zad_4.1/main.go

//		go run zad_6/main.go
