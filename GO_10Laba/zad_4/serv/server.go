package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	secretKey = []byte("your_secret_key") // Замените на ваш секретный ключ
)

// Структура для хранения информации о пользователе
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Структура для JWT токена
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Функция для генерации токена
func generateToken(username string) (string, error) {
	claims := Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// Middleware для проверки токена
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
			c.Abort()
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}
func handler(w http.ResponseWriter, r *http.Request) {
	// var user User

	fmt.Fprintf(w, "Hello, TLS world!")
}
func main() {
	r := gin.Default()

	// Маршрут для входа
	r.POST("/login", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Здесь вы можете добавить проверку пользователя из базы данных
		if user.Username == "admin" && user.Password == "password" { // Пример проверки
			token, err := generateToken(user.Username)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"token": token})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		}
	})

	// Защищенный маршрут
	r.GET("/protected", authMiddleware(), func(c *gin.Context) {
		username := c.MustGet("username").(string)
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Welcome %s!", username)})
	})

	r.Run(":8080") // Запуск сервера на порту 8080
}

// func main() {
// 	// Настройка TLS
// 	tlsConfig := &tls.Config{
// 		MinVersion: tls.VersionTLS13, // Установите минимальную версию TLS
// 	}

// 	// Создание сервера с TLS
// 	server := &http.Server{
// 		Addr:      ":8080",
// 		Handler:   http.HandlerFunc(handler),
// 		TLSConfig: tlsConfig,
// 	}

// 	fmt.Println("Сервер запущен на https://localhost:8080")
// 	// Запуск сервера с сертификатом
// 	if err := server.ListenAndServeTLS("certificate.crt", "private.key"); err != nil {
// 		panic(err)
// 	}
// }

// 	r.RunTLS(":8080", "certificate.crt", "private.key") // Убедитесь, что у вас есть сертификаты
// }

//			go run GO_10Laba/zad_4/serv/server.go
