package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
)

func loadPrivateKeyFromFile(filename string) (*rsa.PrivateKey, error) {
	privPEM, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(privPEM)
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("не удалось расшифровать блок PEM, содержащий закрытый ключ")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func loadPublicKeyFromFile(filename string) (*rsa.PublicKey, error) {
	pubPEM, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(pubPEM)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("не удалось расшифровать блок PEM, содержащий открытый ключ")
	}
	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pub, nil
}

func signMessage(priv *rsa.PrivateKey, message []byte) ([]byte, error) {
	return rsa.SignPKCS1v15(rand.Reader, priv, 0, message)
}

func verifySignature(pub *rsa.PublicKey, message, signature []byte) error {
	return rsa.VerifyPKCS1v15(pub, 0, message, signature)
}

func main() {
	message := []byte("Здравствуйте, это подписанное сообщение!")

	privKey, err := loadPrivateKeyFromFile("private_key.pem")
	if err != nil {
		log.Fatalf("Ошибка при загрузке закрытого ключа: %v", err)
	}

	signature, err := signMessage(privKey, message)
	if err != nil {
		log.Fatalf("Сообщение об ошибке при подписании: %v", err)
	}

	pubKey, err := loadPublicKeyFromFile("public_key.pem")
	if err != nil {
		log.Fatalf("Ошибка при загрузке открытого ключа: %v", err)
	}

	err = verifySignature(pubKey, message, signature)
	if err != nil {
		log.Fatalf("Не удалось подтвердить подпись: %v", err)
	}

	fmt.Println("Подпись успешно подтверждена!")
}

//			go run GO_10Laba/zad_3/2/main.go
