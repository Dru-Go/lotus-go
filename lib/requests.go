package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Request is a struct that holds the request data
type Request struct {
	Method  string
	URL     string
	Header  http.Header
	Payload string
}

// Request sends an HTTP request
func SendHTTPRequest(client http.Client, request Request) ([]byte, error) {
	// Create HTTP request
	httpRequest, err := http.NewRequest(request.Method, request.URL, strings.NewReader(request.Payload))
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
