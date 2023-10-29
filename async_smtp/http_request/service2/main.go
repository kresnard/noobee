package main

import (
	httpclient "async_smtp/http_request/httpClient"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const BASE_URL = "http://localhost:4444/users"

func main() {
	// postUser()
	// getUser()

	// init client
	client := httpclient.NewHttpClient(BASE_URL)

	// proses mengirim request dengan method GET
	resp, err := client.Get("/get")
	if err != nil {
		panic(err)
	}

	log.Println(string(resp))

	reqData := map[string]interface{}{
		"name": "Kresna - Service 2",
	}
	// proses mengirim request dengan method POST
	resp, err = client.Post("/add", reqData)
	if err != nil {
		panic(err)
	}

	log.Println(string(resp))
}

func getUser() {
	// proses melakukan http request dengan method GET
	// bakal jadi : http://localhost:4444/users/get
	resp, err := http.Get(BASE_URL + "/get")
	if err != nil {
		panic(err)
	}

	// close response body jika sudah di baca
	defer resp.Body.Close()

	// proses validasi status code
	// karena dari service1 nge return 200, maka perlu di pastikan bahwa
	// yang di porses hanya yg status code nya 200
	if resp.StatusCode != http.StatusOK {
		log.Println("ada error nih")
		responseBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("error saat parsing response body with error : %v", err.Error())
			return
		}
		log.Println("Response :", string(responseBytes))
		return
	}
	// proses baca response body.
	// jadi hasilnya akan berupa sebuah []byte, yang mana bisa
	// kita jadikan ke sebuah map atau struct nantinya
	// dengan cara melakukan json.Unmarshal()
	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error saat parsing response body with error : %v", err.Error())
		return
	}
	log.Println("Response :", string(responseBytes))
}

func postUser() {
	// generate request data
	reqData := map[string]interface{}{
		"name": "Kresna - Service 2",
	}

	// ubah request data menjadi sebuah []byte
	byteReq, err := json.Marshal(reqData)
	if err != nil {
		panic(err)
	}

	// data yg dikirim via http, adalah sebuah bytes buffer
	// jadi perlu kita ubah dari []byte ke bytes buffer
	buf := bytes.NewBuffer(byteReq)

	/*
			proses http request method "POST"
		 	disini ada 3 parameter, yaitu : URL, ContentType, dan RequestBody
			karena kita akan mengirim sebuah json, maka Content-Type nya adalah : application/json
			untuk request body nya akan diisi dengan bytes buffer tadi
	*/
	resp, err := http.Post(BASE_URL+"/add", "application/json", buf)
	if err != nil {
		panic(err)
	}

	// sisanya sama dengan yg getUser()
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		log.Println("ada error nih")
		responseBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("error saat parsing response body with error : %v", err.Error())
			return
		}
		log.Println("Response :", string(responseBytes))
		return
	}
	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error saat parsing response body with error : %v", err.Error())
		return
	}
	log.Println("Response :", string(responseBytes))
}
