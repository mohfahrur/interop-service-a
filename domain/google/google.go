package google

import (
	"gopkg.in/gomail.v2"
)

type GoogleAgent interface {
	SendEmail(userEmail, Subject, messageBody string) error
}

type GoogleDomain struct {
	Sender   string
	Password string
}

func NewGoogleDomain(sender, password string) *GoogleDomain {

	return &GoogleDomain{
		Sender:   sender,
		Password: password}
}

func (d *GoogleDomain) SendEmail(userEmail, Subject, messageBody string) error {
	// Replace with your Gmail credentials

	message := gomail.NewMessage()
	message.SetHeader("From", d.Sender)
	message.SetHeader("To", userEmail)
	message.SetHeader("Subject", Subject)
	message.SetBody("text/plain", messageBody)

	// Create a new SMTP client
	dialer := gomail.NewDialer("smtp.gmail.com", 587, d.Sender, d.Password)

	// Send the email
	err := dialer.DialAndSend(message)
	if err != nil {
		panic(err)
	}

	println("Email sent successfully!")
	return nil
}
