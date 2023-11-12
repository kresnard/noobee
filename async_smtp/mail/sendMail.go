package mail

import (
	"fmt"
	"net/smtp"
	"strings"

	"gopkg.in/gomail.v2"
)

const (
	// SMTP HOST
	CONFIG_SMTP_HOST = "smtp.gmail.com"
	// SMTP PORT
	CONFIG_SMTP_PORT = 587
	// Sender Name
	CONFIG_SENDER_NAME = "xxxxx <xxxxxx@gmail.com>"

	// AUTH EMAIL FOR GENERATE APP PASSWORD
	CONFIG_AUTH_EMAIL = "xxxxxx@gmail.com"
	// APP PASSWORD THAT HAS BEEN GENERATED
	CONFIG_AUTH_PASSWORD = "xxxx xxxx xxxx xxxx"
)

func SendEmail(to []string, cc []string, subject string, message string) (err error) {
	// pada native package smtp, seluruh header ada di dalam body email
	// jadi perlu di generate di body email nya
	body := "From: " + CONFIG_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(to, ",") + "\n" +
		"Subject: " + subject + "\n\n" + message

	// setup untuk authentication
	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)

	// generate smtpAddress
	// output : smtp.gmail.com:587
	smtpAddress := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	// pross kirim email
	err = smtp.SendMail(smtpAddress, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
	return
}

func SendMailGomail(to []string, cc []string, subject string, message string) (err error) {
	// setup gomail message
	mailer := gomail.NewMessage()

	// settung header from
	mailer.SetHeader("From", CONFIG_SENDER_NAME)

	// setting header to
	mailer.SetHeader("To", to...)

	// setting header CC
	for _, ccEmail := range cc {
		mailer.SetAddressHeader("Cc", ccEmail, "")
	}

	// setting subject
	mailer.SetHeader("Subject", subject)
	// setting body
	// menggunakan body HTML
	mailer.SetBody("text/html", message)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err = dialer.DialAndSend(mailer)

	return
}
