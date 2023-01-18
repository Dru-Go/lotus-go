package lotusgo

import (
	"net/http"
	"time"
)

const BaseURL_V1 = "https://www.uselotus.app/"

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
