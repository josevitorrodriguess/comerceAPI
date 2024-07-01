package services

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/josevitorrodriguess/productsAPI/model"
	"gopkg.in/gomail.v2"
)

func SendMail(reciever string, email model.Email) {
	envPath := filepath.Join("..", ".env")
	fmt.Println("Loading .env file from:", envPath)
	
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	host := "smtp.gmail.com"
	port := 587
	user := os.Getenv("USER")
	password := os.Getenv("EMAIL_PASSWORD")

	dialer := gomail.NewDialer(host, port, user, password)

	msg := gomail.NewMessage()
	msg.SetHeader("From", user)
	msg.SetHeader("To", reciever)
	msg.SetHeader("Subject", email.Subject)
	msg.SetBody("text/plain", email.Body)

	if err := dialer.DialAndSend(msg); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Mensagem enviada!")
}
