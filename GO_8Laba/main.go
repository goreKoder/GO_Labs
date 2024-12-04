package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

const dbURL = "postgres://postgres:123@localhost:5432/postgres" // Замените на свои данные
type User struct {
	First_name  string `json:"first_name"`
	Last_name   string `json:"last_name"`
	Middle_name string `json:"middle_name"`
	Age         int    `json:"age"`
}

func main() {
	// Подключение к базе данных
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer conn.Close(context.Background())

	// Проверка подключения
	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	fmt.Println("Подключение к PostgreSQL успешно установлено!")

	// Создание Gin-сервера
	r := gin.Default()

	// Роуты
	// Получение всех пользователей
	r.GET("/users", func(c *gin.Context) {
		rows, err := conn.Query(context.Background(), "SELECT * FROM public.clients")
		if err != nil {
			log.Printf("Ошибка выполнения запроса: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
			return
		}
		defer rows.Close()
		var users []gin.H
		for rows.Next() {
			var id int
			var first_name string
			var last_name string
			var middle_name string
			var age int
			if err := rows.Scan(&id, &first_name, &last_name, &middle_name, &age); err != nil {
				log.Printf("Ошибка чтения данных: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
				return
			}
			users = append(users, gin.H{"id": id, "username": last_name + " " + first_name + " " + middle_name, "age": age})
		}
		c.JSON(http.StatusOK, gin.H{"users": users})
	})

	// Получение одного пользователя
	r.GET("/users/id/:id", func(c *gin.Context) { // Маршрут с параметром :id
		// Извлекаем параметр id
		id := c.Param("id")
		var sql_com string = "SELECT * FROM public.clients WHERE id = " + id
		rows := conn.QueryRow(context.Background(), sql_com)
		var first_name string  //		имя
		var last_name string   //		фамилия
		var middle_name string //	отчество
		var age int            // возраст
		if err := rows.Scan(&id, &first_name, &last_name, &middle_name, &age); err != nil {
			log.Printf("Ошибка чтения данных: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"id": id, "username": last_name + " " + first_name + " " + middle_name, "age": age})
	})

	// Добавление нового пользователя
	r.POST("/add_user", func(c *gin.Context) {
		var input User

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный запрос"})
			return
		}

		var value_comand string = "INSERT INTO public.clients (first_name, last_name, middle_name, age) VALUES ('" + input.First_name + "','" + input.Last_name + "','" + input.Middle_name + "','" + strconv.Itoa(input.Age) + "')"
		//проверка
		fmt.Println(value_comand)

		_, err := conn.Exec(context.Background(), value_comand)
		if err != nil {
			log.Printf("Ошибка добавления пользователя: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "Пользователь добавлен, а может и нет, фиг знает, проверить надо"})
	})
	//	изменение пользователя
	r.PUT("/change_user/id/:id", func(c *gin.Context) {
		// Извлекаем параметр id
		id := c.Param("id")
		var req User
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		var value_comand string = "UPDATE public.clients \nSET first_name = '" + req.First_name + "', last_name = '" + req.Last_name + "', age = " + strconv.Itoa(req.Age) + ", middle_name = '" + req.Middle_name + "' WHERE id = " + id
		_, err := conn.Exec(context.Background(), value_comand)
		if err != nil {
			log.Printf("Ошибка добавления пользователя: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
		}
		c.JSON(http.StatusOK, gin.H{"status": "Пользователь изменён, а может и нет, фиг знает, проверить надо"})
	})
	r.DELETE("/delete_user/:id", func(c *gin.Context) {
		id := c.Param("id")
		var value_comand string = "DELETE FROM public.clients WHERE id =" + id
		_, err := conn.Exec(context.Background(), value_comand)
		if err != nil {
			log.Printf("Ошибка добавления пользователя: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
		}
		c.JSON(http.StatusOK, gin.H{"status": "Пользователь удален, а может и нет, фиг знает, проверить надо"})
	})
	// Запуск сервера
	r.Run(":8080") // Слушаем порт 8080
}

//			go run main.go
