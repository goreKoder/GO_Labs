package main //		go run GO_7Laba/zad_4/main.go

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Data struct {
	Message string `json:"сообщение какое то хз"`
}

func main() {
	http.HandleFunc("/hello", helloHandler) // Устанавливаем обработчики
	http.HandleFunc("/data", dataHandler)
	err := http.ListenAndServe(":8080", nil) // Запускаем сервер на порту 8080
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
	fmt.Println("Сервер запущен на http://localhost:8080")
}

// Обработчик для корневого маршрута
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
func dataHandler(w http.ResponseWriter, r *http.Request) {
	// Чтение тела запроса
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Выводим полученные данные в консоль
	fmt.Println("Полученные данные:", body)
	fmt.Println("Полученные данные string:", string(body))

	// Отправляем ответ клиенту
	// w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Данные получены успешно.")
}

//		go run GO_7Laba/zad_4/main.go
