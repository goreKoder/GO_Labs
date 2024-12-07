package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

func generateKeyPair(bits int) (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func savePrivateKeyToFile(filename string, priv *rsa.PrivateKey) error {
	privBytes := x509.MarshalPKCS1PrivateKey(priv)
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privBytes,
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return pem.Encode(file, block)
}

func savePublicKeyToFile(filename string, pub *rsa.PublicKey) error {
	pubBytes := x509.MarshalPKCS1PublicKey(pub)
	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubBytes,
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return pem.Encode(file, block)
}

func main() {
	privateKey, err := generateKeyPair(2048)
	if err != nil {
		log.Fatalf("Error generating key pair: %v", err)
	}

	err = savePrivateKeyToFile("private_key.pem", privateKey)
	if err != nil {
		log.Fatalf("Error saving private key: %v", err)
	}

	err = savePublicKeyToFile("public_key.pem", &privateKey.PublicKey)
	if err != nil {
		log.Fatalf("Error saving public key: %v", err)
	}

	fmt.Println("Keys generated and saved to files.")
}

//		go run GO_10Laba/zad_3/1/main.go
