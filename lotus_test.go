package lotusgo

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type BasicClient struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
	debug      bool
}

var TestURL_V1 = "https://api.uselotus.io"
var TestAPI_KEY = "ZtECIKzt.r2NGXzU8lqsghdqNABreakedpS5k3Ju1"

var mockClient = BasicClient{
	BaseURL: TestURL_V1,
	apiKey:  TestAPI_KEY,
	HTTPClient: &http.Client{
		Timeout: time.Minute,
	},
	debug: true,
}

func TestClient_ListCustomers(t *testing.T) {

	tests := []struct {
		name   string
		fields BasicClient
		want   []CustomerResponse
	}{
		{
			name:   "Basic Test",
			fields: mockClient,
			want:   []CustomerResponse{},
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
			got, err := client.ListCustomers()
			assert.Nil(t, err)
			fmt.Printf("%+v\n", got)
			assert.NotNil(t, got)
		})
	}
}

func TestClient_GetCustomer(t *testing.T) {
	type args struct {
		message CustomerDetailsParams
	}
	tests := []struct {
		name    string
		fields  BasicClient
		args    args
		want    CustomerResponse
		wantErr bool
	}{
		{
			name:   "Basic Test",
			fields: mockClient,
			args: args{
				message: CustomerDetailsParams{
					CustomerId: "1234567",
				},
			},
			want: CustomerResponse{},
		},
		{
			name:   "Existing Record Test",
			fields: mockClient,
			args: args{
				message: CustomerDetailsParams{
					CustomerId: "123456",
				},
			},
			want: CustomerResponse{
				CustomerId:       "123456",
				Email:            "silco2dev@gmail.com",
				CustomerName:     "Silco",
				Invoices:         []Invoice{},
				TotalAmountDue:   0,
				Subscriptions:    []CreateSubscription{},
				Integrations:     Integrations{},
				DefaultCurrency:  DefaultCurrency{},
				HasPaymentMethod: false,
				PaymentProvider:  "",
				Address:          "",
				TaxRate:          int16(0),
			},
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
			got, err := client.GetCustomer(tt.args.message)
			assert.Nil(t, err)
			assert.NotNil(t, got)
		})
	}
}

func TestClient_CreateCustomer(t *testing.T) {
	type args struct {
		message CreateCustomerParams
	}
	tests := []struct {
		name    string
		fields  BasicClient
		args    args
		want    CustomerResponse
		wantErr bool
	}{
		{
			name:   "Basic Create Customer Test",
			fields: mockClient,
			args:   args{},
			want:   CustomerResponse{},
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
			got, err := client.CreateCustomer(tt.args.message)
			assert.Nil(t, err)
			assert.NotNil(t, got)
		})
	}
}
