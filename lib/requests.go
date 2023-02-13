package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Request is a struct that holds the request data
type Request struct {
	Method  string
	URL     string
	Header  http.Header
	Payload interface{}
}

// SendHTTPRequest sends an HTTP request
func SendHTTPRequest(client http.Client, request Request) ([]byte, error) {
	// Marshal request payload
	var requestBody []byte
	if request.Payload != nil {
		var err error
		requestBody, err = json.Marshal(request.Payload)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	// Create HTTP request
	httpRequest, err := http.NewRequest(request.Method, request.URL, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	httpRequest.Header = request.Header

	// Send HTTP request
	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer httpResponse.Body.Close()

	// Read HTTP response
	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return responseBody, nil
}
