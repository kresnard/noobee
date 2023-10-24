package main

import "strings"

const (
	// SMTP HOST
	CONFIG_SMTP_HOST = "smtp.gmail.com"
	// SMTP PORT
	CONFIG_SMTP_PORT = 587
	// Sender Name
	CONFIG_SENDER_NAME = "Kresna <kresnard@gmail.com>"

	// AUTH EAMAIL FOR GENERATE APP PASSWORD
	CONFIG_AUTH_EMAIL = "kresnard@gmail.com"
	// APP PASSWORD THAT HAS BEEN GENERATED
	CONFIG_AUTH_PASSWORD = "bohm nrnd xyhe qhkg"
)

func sendEmail(to []string, cc []string, subject string, message string) (err error) {
	body := "From: " + CONFIG_AUTH_EMAIL + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(to, ",") + "\n" +
		"Subject: " + subject + "\n\n" + message

	return
}
