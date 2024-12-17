// package main

// import (
// 	"crypto/tls"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// func main() {
// 	// Создание TLS-клиента
// 	tlsConfig := &tls.Config{
// 		InsecureSkipVerify: true, // Для тестирования с самоподписанным сертификатом
// 	}

// 	// Создание HTTP-клиента с TLS
// 	transport := &http.Transport{
// 		TLSClientConfig: tlsConfig,
// 	}
// 	client := &http.Client{Transport: transport}

// 	// Отправка GET-запроса
// 	resp, err := client.Get("https://localhost:8080")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()

// 	// Чтение ответа
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		panic(err)
// 	}

//		fmt.Println(string(body))
//	}
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func main() {
	// Вход пользователя
	user := User{
		Username: "admin",
		Password: "password",
	}

	// Отправка запроса на вход
	userData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshaling user data:", err)
		return
	}

	loginResp, err := http.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer(userData))
	if err != nil {
		fmt.Println("Error logging in:", err)
		return
	}
	defer loginResp.Body.Close()

	if loginResp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(loginResp.Body)
		fmt.Println("Login failed:", string(body))
		return
	}

	var tokenResponse TokenResponse
	json.NewDecoder(loginResp.Body).Decode(&tokenResponse)

	fmt.Println("Token received:", tokenResponse.Token)

	// Доступ к защищенному маршруту
	req, err := http.NewRequest("GET", "http://localhost:8080/protected", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Authorization", tokenResponse.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error accessing protected route:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("Access denied:", string(body))
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Protected route response:", string(body))
}

//			go run cli/client.go
//			go run client.go
// 			go run GO_10Laba/zad_4/cli/client.go
