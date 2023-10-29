package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/gomail.v2"
)

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

type EmailRequest struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Cc      []string `json:"cc"`
	Subject string   `json:"subject"`
	Message string   `json:"message"`
	Type    string   `json:"type"`
}

type EmailResponse struct {
	Success bool   `json:"successs"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func main() {
	http.HandleFunc("/send", sendEmailHandler)
	http.ListenAndServe(":8000", nil)
	log.Println("service running in port 8000")
}

func sendEmailHandler(w http.ResponseWriter, r *http.Request) {
	var req EmailRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := SendMailGomail(req.To, req.Cc, req.From, req.Subject, req.Message)
	if err != nil {
		response := EmailResponse{
			Success: false,
			Message: "Email failed to send",
			Error:   err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	response := EmailResponse{
		Success: true,
		Message: "Email sent successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	json.NewEncoder(w).Encode(response)
}

func SendMailGomail(to, cc []string, from, subject, message string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", to...)
	for _, ccEmail := range cc {
		mailer.SetAddressHeader("Cc", ccEmail, "")
	}
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", message)

	dialer := gomail.NewDialer(CONFIG_SMTP_HOST, CONFIG_SMTP_PORT, CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD)
	if err := dialer.DialAndSend(mailer); err != nil {
		fmt.Println("Failed to send email:", err)
		return err
	}
	return nil
}
