package google

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"

	"golang.org/x/oauth2/google"
	gmail "google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type GoogleAgent interface {
	CreateMessage(sender, recipient, subject, body string) *gmail.Message
	SendEmail(userID string, message *gmail.Message) error
}

type GoogleDomain struct {
	gmailService *gmail.Service
}

func NewGoogleDomain(credentialsFile []byte) *GoogleDomain {
	creds, err := google.CredentialsFromJSON(context.Background(), credentialsFile, gmail.MailGoogleComScope)
	if err != nil {
		log.Fatalf("failed to read service account credentials: %v", err)
	}

	service, err := gmail.NewService(context.Background(), option.WithCredentials(creds))
	if err != nil {
		log.Fatalf("failed to create Gmail service: %v", err)
	}
	return &GoogleDomain{gmailService: service}
}

func (d *GoogleDomain) CreateMessage(sender, recipient, subject, body string) *gmail.Message {
	message := &gmail.Message{
		Raw: base64.RawURLEncoding.EncodeToString([]byte(
			fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", sender, recipient, subject, body))),
	}

	return message
}

func (d *GoogleDomain) SendEmail(userID string, message *gmail.Message) error {
	_, err := d.gmailService.Users.Messages.Send(userID, message).Do()
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
