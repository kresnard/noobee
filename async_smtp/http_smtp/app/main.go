package main

import (
	"bytes"
	"encoding/json"
	"net/http"
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
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func main() {
	http.HandleFunc("/send", sendEmailHandler)
	http.ListenAndServe(":8001", nil)
}

func sendEmailHandler(w http.ResponseWriter, r *http.Request) {
	var request EmailRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mailServiceURL := "http://localhost:8000/send"
	mailServiceRequest, _ := json.Marshal(request)

	resp, err := http.Post(mailServiceURL, "application/json", bytes.NewBuffer(mailServiceRequest))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var response EmailResponse
	decoder = json.NewDecoder(resp.Body)
	if err := decoder.Decode(&response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	json.NewEncoder(w).Encode(response)
}
