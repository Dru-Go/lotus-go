package lotusgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/lotus-go/lib"
)

// "github.com/lotus-go/lib"

var BaseURL_V1 = "https://api.uselotus.io" // process env

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

func (client *Client) Ping() (PingResponse, error) {
	endpoint := BaseURL_V1 + PING
	request := lib.Request{Method: "GET", URL: endpoint, Header: http.Header{}, Payload: ""}
	request.Header.Add("X-API-KEY", client.apiKey)

	body, err := lib.SendHTTPRequest(*client.HTTPClient, request)
	if err != nil {
		return PingResponse{}, err
	}

	var response PingResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return PingResponse{}, err
	}

	return response, nil
}

func (client *Client) ListCustomers() ([]CustomerResponse, error) {
	endpoint := BaseURL_V1 + GET_CUSTOMERS
	request := lib.Request{Method: "GET", URL: endpoint, Header: http.Header{}, Payload: ""}
	request.Header.Add("X-API-KEY", client.apiKey)

	body, err := lib.SendHTTPRequest(*client.HTTPClient, request)
	if err != nil {
		return []CustomerResponse{}, err
	}

	var response []CustomerResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return []CustomerResponse{}, err
	}

	return response, nil
}

func (client *Client) GetCustomer(message CustomerDetailsParams) (CustomerResponse, error) {
	endpoint := fmt.Sprint(BaseURL_V1, GET_CUSTOMERS, message.CustomerId)
	request := lib.Request{Method: "GET", URL: endpoint, Header: http.Header{}, Payload: ""}
	request.Header.Add("X-API-KEY", client.apiKey)

	body, err := lib.SendHTTPRequest(*client.HTTPClient, request)
	if err != nil {
		return CustomerResponse{}, err
	}

	var response CustomerResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return CustomerResponse{}, err
	}

	return response, nil
}

func (client *Client) CreateCustomer(params CreateCustomerParams) (CustomerResponse, error) {
	endpoint := fmt.Sprint(BaseURL_V1, CREATE_CUSTOMERS)
	formatted, err := query.Values(params)
	if err != nil {
		return CustomerResponse{}, err
	}
	request := lib.Request{Method: "POST", URL: endpoint, Header: http.Header{}, Payload: formatted.Encode()}
	request.Header.Add("X-API-KEY", client.apiKey)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	body, err := lib.SendHTTPRequest(*client.HTTPClient, request)
	if err != nil {
		return CustomerResponse{}, err
	}
	var response CustomerResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return CustomerResponse{}, err
	}

	return response, nil
}

func (client *Client) TrackEvent(params TrackEventParams) (TrackEventResponse, error) {
	endpoint := fmt.Sprint(BaseURL_V1, TRACK_EVENT)
	// JSON stringify
	val, err := json.Marshal(params)
	if err != nil {
		return TrackEventResponse{}, err
	}
	// fmt.Printf("%+v\n", string(val))

	request := lib.Request{Method: "POST", URL: endpoint, Header: http.Header{}, Payload: string(val)}
	request.Header.Add("X-API-KEY", client.apiKey)
	request.Header.Add("Content-Type", "application/json")

	body, err := lib.SendHTTPRequest(*client.HTTPClient, request)
	if err != nil {
		return TrackEventResponse{}, err
	}
	var response TrackEventResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return TrackEventResponse{}, err
	}

	return response, nil
}
