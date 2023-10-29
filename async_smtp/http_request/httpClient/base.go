package httpclient

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// objek untuk melakukan http request
type HTTPClient struct {
	baseUrl string
	client  http.Client
}

// fungsi untuk init object
func NewHttpClient(baseUrl string) HTTPClient {
	client := http.Client{}
	return HTTPClient{
		client:  client,
		baseUrl: baseUrl,
	}
}

/*
membutuhkan 3 parameter, yaitu method, url dan request body.
*/
func (h HTTPClient) run(method string, url string, reqBody []byte) (response []byte, err error) {
	// request body akan dibungkus dengan bytes.NewBuffer
	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	// proses baca response dan menjadikannya []byte
	response, err = io.ReadAll(resp.Body)
	return
}

/*
membutuhkan 2 parameter, yaitu endpoint dan request body.
Fungsi ini akan melakukan marshal terhadap request body dan mengirimnya ke method run()
*/
func (h HTTPClient) Post(endpoint string, requestBody interface{}) (response []byte, err error) {
	// proses membuat full URL
	url := h.baseUrl + endpoint

	// mengubah interface (object) menjadi []byte
	reqByte, err := json.Marshal(requestBody)
	if err != nil {
		return
	}

	// mengirim data ke method run()
	return h.run(http.MethodPost, url, reqByte)
}

/*
membutuhkan 1 parameter, yaitu endpoint.
*/
func (h HTTPClient) Get(endpoint string) (response []byte, err error) {
	url := h.baseUrl + endpoint
	return h.run(http.MethodGet, url, nil)
}

func (h HTTPClient) Put(endpoint string, requestBody interface{}) (response []byte, err error) {
	// proses membuat full URL
	url := h.baseUrl + endpoint

	// mengubah interface (object) menjadi []byte
	reqByte, err := json.Marshal(requestBody)
	if err != nil {
		return
	}

	// mengirim data ke method run()
	return h.run(http.MethodPut, url, reqByte)
}
