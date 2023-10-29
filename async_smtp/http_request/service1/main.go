package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// define port application
const (
	APP_PORT = ":4444"
)

// setup object user
type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// menampung list user
var users = []User{
	{
		Id:   1,
		Name: "Kresna",
	},
	{
		Id:   2,
		Name: "Rangga",
	},
}

func main() {
	// GET /users/get
	http.HandleFunc("/users/get", getUser)

	// POST /users/add
	http.HandleFunc("/users/add", addUser)

	log.Println("server running at port", APP_PORT)
	http.ListenAndServe(APP_PORT, nil)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// validasi hanya method GET yang boleh hit endpoint ini
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "method not allowed",
		})
	}

	json.NewEncoder(w).Encode(users)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// makesure method yang digunakan adalah POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "method not allowed",
		})
	}

	var req = User{}
	// proses parsing request dari request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
	}

	// generate id
	req.Id = len(users) + 1

	// proses menambahkan user
	users = append(users, req)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(users)
}
