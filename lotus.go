package lotusgo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/lotus-go/lib"
)

// "github.com/lotus-go/lib"

const BaseURL_V1 = "https://api.uselotus.io" // process env

type Client struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
	debug      bool
}

func NewClient(apiKey string) *Client {
	return &Client{
		BaseURL: BaseURL_V1,
		apiKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
		debug: false,
	}
}

func NewClientWithTimeOut(apiKey string, timeOut time.Duration) *Client {
	return &Client{
		BaseURL: BaseURL_V1,
		apiKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: timeOut,
		},
		debug: false,
	}
}

func (client *Client) ListCustomers() ([]ListCustomerResponse, error) {
	// Request HTTP with timeout
	req, _ := http.NewRequest("GET", BaseURL_V1+GET_CUSTOMERS, nil)
	req.Header.Add("X-API-KEY", fmt.Sprint(client.apiKey))

	// Make the request to the API
	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return []ListCustomerResponse{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		panic(err.Error())
	}

	log.Printf("body = %v", string(body))

	var response []ListCustomerResponse

	// Unmarshal the JSON response into a slice of ListCustomerResponse
	err3 := json.Unmarshal(body, &response)
	if err3 != nil {
		log.Printf("error = %v", err3)
		return []ListCustomerResponse{}, err3
	}

	log.Printf("s = %v", response)

	return response, nil
}

func (client *Client) ListCustomersV2() ([]ListCustomerResponse, error) {
	endpoint := BaseURL_V1 + GET_CUSTOMERS
	cli := NewClient(client.apiKey)
	request := lib.Request{Method: "GET", URL: endpoint, Header: http.Header{}, Payload: nil}
	request.Header.Add("X-API-KEY", fmt.Sprint(client.apiKey))

	body, err := lib.SendHTTPRequest(*cli.HTTPClient, request)
	if err != nil {
		fmt.Println(err)
		return []ListCustomerResponse{}, err
	}

	var response []ListCustomerResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return []ListCustomerResponse{}, err
	}

	return response, nil
}
