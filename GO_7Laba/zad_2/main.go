package main //		go run GO_7Laba/zad_2/main.go

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000") // устанавили соединение с клиентом
	if err != nil {
		fmt.Println("Ерор:", err.Error())
		return
	}
	defer conn.Close()

	fmt.Println("подрубился по TCP емае")
	message := "здарова"
	fmt.Scan(&message)
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending data:", err.Error())
		return
	}
	fmt.Println("Отправили сообщение")

}

//		go run GO_7Laba/zad_2/main.go
//		go run zad_2/main.go
