package service

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendEmail(hash string) {
	from := "test@mail.ru"

	user := "test@mail.ru"
	password := "11111"

	to := []string{
		"test@mail.ru",
	}

	addr := "smtp.yandex.ru:465"
	host := "smtp.yandex.ru"

	msg := []byte("From: test@mail.ru\r\n" +
		"To: test@mail.ru\r\n" +
		"Subject: Test mail\r\n\r\n" +
		"Email body\r\n")

	auth := smtp.PlainAuth("", user, password, host)

	err := smtp.SendMail(addr, auth, from, to, msg)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent successfully")

}
