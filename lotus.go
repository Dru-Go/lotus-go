package lotusgo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const BaseURL_V1 = "https://www.uselotus.app"

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

func (client *Client) ListCustomers() ([]ListCustomerResponse, error) {
	// Request HTTP with timeout
	req, _ := http.NewRequest("GET", BaseURL_V1+GET_CUSTOMERS, nil)
	req.Header.Add("X-API-KEY", client.apiKey)

	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return []ListCustomerResponse{}, err
	}
	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		panic(err.Error())
	}

	log.Printf("body = %v", string(body))

	var response []ListCustomerResponse

	err3 := json.Unmarshal(body, &response)
	if err3 != nil {
		log.Printf("error = %v", err3)
		return []ListCustomerResponse{}, err3
	}

	log.Printf("s = %v", response)

	return response, nil
}
