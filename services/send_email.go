package services

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"phone_email/settings"
)

func SendEmail(numberData string, goalEmail string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", settings.EmailUsername)
	m.SetHeader("To", goalEmail)
	m.SetHeader("Subject", settings.EmailHeaderSubject)
	m.SetBody("text/html", fmt.Sprintf("Numbers: %s", numberData))

	d := gomail.NewDialer(settings.EmailHost, settings.EmailPort, settings.EmailUsername, settings.EmailPassword)

	return d.DialAndSend(m)
}
