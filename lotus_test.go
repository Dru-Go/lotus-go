package lotusgo

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const TestURL_V1 = "https://api.uselotus.io/"
const TestAPI_KEY = "ZtECIKzt.r2NGXzU8lqsghdqNABreakedpS5k3Ju1"

func TestClient_ListCustomers(t *testing.T) {
	type fields struct {
		BaseURL    string
		apiKey     string
		HTTPClient *http.Client
		debug      bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []ListCustomerResponse
	}{
		{
			name: "Basic Test",
			fields: fields{
				BaseURL: TestURL_V1,
				apiKey:  TestAPI_KEY,
				HTTPClient: &http.Client{
					Timeout: time.Minute,
				},
				debug: true,
			},
			want: []ListCustomerResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				BaseURL:    tt.fields.BaseURL,
				apiKey:     tt.fields.apiKey,
				HTTPClient: tt.fields.HTTPClient,
				debug:      tt.fields.debug,
			}
			got, err := client.ListCustomersV2()
			assert.Nil(t, err)
			assert.NotNil(t, got)
		})
	}
}
