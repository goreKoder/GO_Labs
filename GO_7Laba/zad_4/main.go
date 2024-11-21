package main //		go run GO_7Laba/zad_4/main.go

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Data struct {
	Message string `json:"сообщение какое то хз"`
}

func main() {
	http.HandleFunc("/hello", middleware(helloHandler)) // Устанавливаем обработчики
	http.HandleFunc("/data", middleware(dataHandler))
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
	fmt.Println()
}
func middleware(next http.HandlerFunc) http.HandlerFunc { //карабли логировали логировали да не вылогировали
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now() // Запоминаем время начала обработки запроса
		next(w, r)
		// log.Println("логирую в ретёрне")
		duration := time.Since(startTime) // Вычисляем время выполнения
		method := r.Method                // Получаем HTTP метод
		url := r.URL.Path                 // Получаем URL
		log.Printf("Метод: %s, URL: %s, время выполнения: %d\n", method, url, duration.Milliseconds())
	}
}

//		go run GO_7Laba/zad_4/main.go
