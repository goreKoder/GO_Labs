package main //      go run GO_7Laba/zad_1/main.go

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000") // приказываю слушать порт
	if err != nil {
		fmt.Println("Error while listening:", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("Порт слушается емае, localhost:8000")

	for { //цикл бесконечностей емае
		conn, err := listener.Accept() // принимаю новых клиентов (вроде)
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}

		go handleClient(conn) //горутина чтобы чуть что всех закрыть нафиг
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	clientAddr := conn.RemoteAddr().String() //Я получил твой аднес, АХАХАХААХХАХАХАХАХАХХАХАХ!
	fmt.Println("Адрес клиента:", clientAddr)

	for {
		// считываем полученные в запросе данные
		input := make([]byte, (1024 * 4))
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("ошибка:", err)
			break
		}
		source := string(input[0:n])
		fmt.Println("Сообщение от клиента: ", source)
	}
	// fmt.Println("Client disconnected:", clientAddr)
}

//      go run GO_7Laba/zad_1/main.go
