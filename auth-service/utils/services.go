package utils_auth_server

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"log"

	configl_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/config"
	"github.com/wneessen/go-mail"
	go_mail "gopkg.in/mail.v2"
)

func SendVerificationEmail(recipientEmail, verificationToken string, credentials configl_auth_server.Mail) error {
	tmpl, err := template.ParseFiles("template/mail.html")
	if err != nil {
		log.Fatalf("failed to parse email template: %s", err)
	}

	data := struct {
		VerificationURL string
	}{
		VerificationURL: fmt.Sprintf("http://hyperhive.vajid.tech:8000/user/verify?email=%s&token=%s", recipientEmail, verificationToken),
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

func SendOtp(toEmail, otp string, TemperveryToken string, credentials configl_auth_server.Mail) error {
	// type data struct{
	// 	OTP string
	// }
	otpData := struct {
		OTP string
	}{
		OTP: otp,
	}

	// t := template.New("template/otp.html")

	var err error
	t, err := template.ParseFiles("template/otp.html")
	if err != nil {
		return err
	}

	var tpl bytes.Buffer
	if err = t.Execute(&tpl, otpData); err != nil {
		fmt.Println("--",tpl.String())
		return err
	}

	result := tpl.String()

	m := go_mail.NewMessage()

	m.SetHeader("From", credentials.From)

	m.SetHeader("To", toEmail)

	m.SetHeader("Subject", "Otp Verification from HyperHive")

	m.SetBody("text/html", result)

	// Settings for SMTP server
	d := go_mail.NewDialer("smtp.gmail.com", 587, credentials.From, credentials.SecretKey)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return nil
}
