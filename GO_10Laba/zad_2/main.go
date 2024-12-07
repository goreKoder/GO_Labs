package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

// encrypt encrypts the given plaintext using AES-256-GCM.
func encrypt(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

// decrypt decrypts the given ciphertext using AES-256-GCM.
func decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func main() {
	var plaintext, keyStr string

	fmt.Print("Введите строку для шифрования: ")
	fmt.Scanln(&plaintext)

	fmt.Print("Введите секретный ключ (32 символа): ")
	fmt.Scanln(&keyStr)

	//Проверка длины ключа
	if len(keyStr) != 32 {
		fmt.Println("Ошибка: Длина ключа должна быть 32 символа (256 бит).")
		return
	}

	key, err := hex.DecodeString(keyStr)
	if err != nil {
		fmt.Printf("Ошибка при декодировании ключа: %v\n", err)
		return
	}

	ciphertext, err := encrypt([]byte(plaintext), key)
	if err != nil {
		fmt.Printf("Ошибка при шифровании: %v\n", err)
		return
	}

	fmt.Printf("Зашифрованный текст (hex): %x\n", ciphertext)

	decrypted, err := decrypt(ciphertext, key)
	if err != nil {
		fmt.Printf("Ошибка при расшифровании: %v\n", err)
		return
	}

	fmt.Printf("Расшифрованный текст: %s\n", decrypted)
}

// 			go run GO_10Laba/zad_2/main.go
//			12345678123456781234567812345678
