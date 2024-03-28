package utils_auth_server

import (
	"bytes"
	"fmt"
	"html/template"
	"log"

	configl_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/config"
	"github.com/wneessen/go-mail"
)

func SendVerificationEmail(recipientEmail, verificationToken string, credentials configl_auth_server.Mail) error {
	tmpl, err := template.ParseFiles("template/mail.html")
	if err != nil {
		log.Fatalf("failed to parse email template: %s", err)
	}

	data := struct {
		VerificationURL string
	}{
		VerificationURL: fmt.Sprintf("http://hyperhive.vajid.tech:8080/user/verify?email=%s&token=%s", recipientEmail, verificationToken),
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		log.Fatalf("failed to execute email template: %s", err)
	}

	// Define email message
	m := mail.NewMsg()
	m.From(credentials.From)
	m.To(recipientEmail)
	m.Subject("Verification Mail From HiperHive!")
	m.SetBodyString(mail.TypeTextHTML, body.String())

	// Set the email body
	// m.SetBodyString(mail.TypeTextPlain, body)

	c, err := mail.NewClient("smtp.gmail.com", mail.WithPort(587), mail.WithSMTPAuth(mail.SMTPAuthPlain), mail.WithUsername(credentials.From), mail.WithPassword(credentials.SecretKey))
	if err != nil {
		log.Fatalf("failed to create mail client: %s", err)
	}

	if err := c.DialAndSend(m); err != nil {
		log.Fatalf("failed to send mail: %s", err)
	}

	return nil
}
