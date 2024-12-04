package main //		go run GO_9Laba/main.go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	First_name  string `json:"first_name"`
	Last_name   string `json:"last_name"`
	Middle_name string `json:"middle_name"`
	Age         int    `json:"age"`
}

// функция для запуска функций
func variants() {
	var operation_num int
	fmt.Scan(&operation_num)
	switch operation_num {
	case 0:
		return
	case 1:
		Get_all()
		variants()
	case 2:
		Get_one()
		variants()
	case 3:
		Post_create_user()
		variants()
	case 4:
		Put_change_user()
		variants()
	case 5:
		Delete_user()
		variants()
	}

}
func main() {
	fmt.Println("Добро пожаловать, введите предпочтительную операцию: ")
	fmt.Println("Вот их список: \n1. Вывод всех пользователей \n2. Вывод конкретного пользователя\n3. Добавление пользователя\n4. Изменение пользователя\n5. Удаление пользователя")
	//		запускаем бесконечные запросы
	variants()
}

// запрос на получение всех пользователей
func Get_all() {
	resp, err := http.Get("http://localhost:8080/users")
	if err != nil {
		log.Println("Ошибка при создании запроса", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Жопа с получением ответа", err)
		return
	}
	fmt.Println("Ответ от сервера (GET/users):", string(body))
}

// запрос на получение одного пользователя
func Get_one() {
	var id int
	fmt.Println("Введи ID: ")
	fmt.Scan(&id)
	zapros_text := "http://localhost:8080/users/id/" + strconv.Itoa(id)
	resp, err := http.Get(zapros_text)
	if err != nil {
		log.Println("Ошибка при создании запроса", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Жопа с получением ответа", err)
		return
	}
	fmt.Println("Ответ от сервера (GET/users/id):", string(body))
}

// добавление пользователя
func Post_create_user() {
	// var id int
	var first_name string  //	имя
	var last_name string   //	фамилия
	var middle_name string //	отчество
	var age int            //	возрост

	fmt.Println("Введите: имя, фамилия, отчество, возраст")
	fmt.Scan(&first_name, &last_name, &middle_name, &age)
	new_user := User{
		First_name:  first_name,  //	имя
		Last_name:   last_name,   //	фамилия
		Middle_name: middle_name, //	отчество
		Age:         age,         //	возрост
	}
	// Преобразование структуры в JSON
	jsonData, err := json.Marshal(new_user)
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return
	}
	requestBody := []byte(jsonData)
	resp, err := http.Post("http://localhost:8080/add_user", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Panicln("ошибка в отправке Post запроса на сервер", err)
		return
	}
	defer resp.Body.Close()
	// Чтение и вывод ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Ошибка при чтении ответа:", err)
		return
	}
	fmt.Println("Ответ от сервера (POST/add_user):", string(body))
}
func Put_change_user() {
	var first_name string  //	имя
	var last_name string   //	фамилия
	var middle_name string //	отчество
	var age int            //	возрост
	var id int
	fmt.Println("Введите: имя, фамилия, отчество, возраст")
	fmt.Scan(&first_name, &last_name, &middle_name, &age)
	fmt.Println("Введите ID пользователя которого надо подменить: ")
	fmt.Scan(&id)
	zapros_text := "http://localhost:8080/change_user/id/" + strconv.Itoa(id)
	new_user := User{
		First_name:  first_name,  //	имя
		Last_name:   last_name,   //	фамилия
		Middle_name: middle_name, //	отчество
		Age:         age,         //	возрост
	}
	// Кодирование данных в JSON
	jsonData, err := json.Marshal(new_user)
	if err != nil {
		log.Fatalf("Ошибка кодирования JSON: %v", err)
	}
	// Создание PUT-запроса
	req, err := http.NewRequest(http.MethodPut, zapros_text, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Ошибка создания PUT-запроса: %v", err)
	}

	// Установка заголовков
	req.Header.Set("Content-Type", "application/json")

	// Отправка запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка отправки PUT-запроса: %v", err)
	}
	// Чтение и вывод ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Ошибка при чтении ответа:", err)
		return
	}
	fmt.Println("Ответ от сервера (PUT/add_user):", string(body))
	defer resp.Body.Close()
}
func Delete_user() {
	var id int
	fmt.Println("Введи id чувака для удаления: ")
	fmt.Scan(&id)
	zapros_text := "http://localhost:8080/delete_user/" + strconv.Itoa(id)
	req, err := http.NewRequest(http.MethodDelete, zapros_text, nil)
	// Добавляем заголовки, если нужно
	req.Header.Set("Authorization", "Bearer your_token_here") // Пример: токен авторизации

	// Отправляем запрос с клиентом
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Ошибка отправки DELETE-запроса: %v\n", err)
		return
	}
	// Чтение и вывод ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Ошибка при чтении ответа:", err)
		return
	}
	fmt.Println("Ответ от сервера (DELERE/add_user):", string(body))
	defer resp.Body.Close()
}

//		go run GO_9Laba/main.go
