package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {
	//задание 1
	var choice int
	fmt.Println("Выберите какой хеш функцией хотите воспользоваться: 1. MD5, 2. SHA-256, 3. SHA-512")
	fmt.Scan(&choice)
	var pasvord string
	var hash string
	fmt.Print("Введите что-то: ")
	fmt.Scan(&pasvord)
	fmt.Scan(&hash)
	switch choice {
	case 1:
		fmt.Printf("%x", md5.Sum([]byte(pasvord)))
	case 2:
		fmt.Printf("%x", sha256.Sum256([]byte(pasvord)))
	case 3:
		fmt.Printf("%x", sha512.Sum512([]byte(pasvord)))
	}

}

//		go run GO_10Laba/zad_1/main.go
