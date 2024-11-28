package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	// "github.com/gorilla/websocket"
)

// Создаем структуру для хранения информации о клиенте
type Client struct {
	conn *websocket.Conn
	send chan []byte
}

// Хранение всех подключенных клиентов
var clients = make(map[*Client]bool)
var clientsMutex sync.Mutex

// Устанавливаем веб-сокет
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Разрешаем все подключения
	},
}

// Обрабатываем входящие сообщения от клиентов
func handleMessages(client *Client) {
	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			fmt.Println("Ошибка при чтении сообщения:", err)
			break
		}
		broadcastMessage(message)
	}
	client.conn.Close()
}

// Рассылаем сообщения всем подключенным клиентам
func broadcastMessage(message []byte) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	for client := range clients {
		select {
		case client.send <- message:
		default:
			close(client.send)
			delete(clients, client)
		}
	}
}

// Обрабатываем подключения клиентов
func handleClientConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Ошибка при подключении:", err)
		return
	}
	client := &Client{conn: conn, send: make(chan []byte)}
	clientsMutex.Lock()
	clients[client] = true
	clientsMutex.Unlock()

	go handleMessages(client)

	// Отправляем сообщения клиенту
	for message := range client.send {
		if err := client.conn.WriteMessage(websocket.TextMessage, message); err != nil {
			fmt.Println("Ошибка при отправке сообщения:", err)
			break
		}
	}
}

// Основная функция
func main() {
	http.HandleFunc("/ws", handleClientConnection)

	fmt.Println("Сервер запущен на :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}

//				go mod init github.com/gorilla/websocket
//				go get github.com/gorilla/websocket
//				go get -u github.com/gorilla/websocket

//		go run main.go
//		go run GO_7Laba/zad_6/main.go
//		go clean -modcache
//		go mod tidy
