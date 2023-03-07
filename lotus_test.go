package lotusgo

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
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
	customerId := uuid.NewString()
	type args struct {
		params CreateCustomerParams
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
			args: args{
				params: CreateCustomerParams{
					CustomerId:   customerId,
					Email:        "derraaadugna2@gmail.com",
					CustomerName: "Dre3",
				},
			},
			want: CustomerResponse{
				CustomerId:    customerId,
				Email:         "derraaadugna2@gmail.com",
				CustomerName:  "Dre3",
				Invoices:      []Invoice{},
				Subscriptions: []CreateSubscription{},
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
			got, err := client.CreateCustomer(tt.args.params)
			assert.Nil(t, err)
			assert.NotNil(t, got)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestClient_Ping(t *testing.T) {
	tests := []struct {
		name    string
		fields  BasicClient
		want    bool
		wantErr bool
	}{
		{
			name:   "Basic Ping Request",
			fields: mockClient,
			want:   true,
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
			got, err := client.Ping()
			assert.Nil(t, err)
			assert.NotNil(t, got)
			assert.NotNil(t, got.OrganizationId)
		})
	}
}

func TestClient_TrackEvent(t *testing.T) {
	type args struct {
		params TrackEventParams
	}
	tests := []struct {
		name   string
		fields BasicClient
		args   args
		want   TrackEventResponse
	}{
		{
			name:   "TrackEvent Test",
			fields: mockClient,
			args: args{
				params: TrackEventParams{
					Batch: []TrackEventEntity{
						{
							CustomerId:  "cust_ddfae050b12c4c3a9e092917299296c4",
							EventName:   "api_post",
							ImpotencyId: time.Now().Format(time.RFC3339),
							Properties: TrackEventProperty{
								ShardId:   "2",
								ShardType: "professional",
								Change:    int32(4),
							},
							TimeCreated: time.Now(),
						},
					},
				},
			},
			want: TrackEventResponse{
				Success: "all",
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
			got, err := client.TrackEvent(tt.args.params)
			assert.Nil(t, err)
			assert.NotNil(t, got)
			assert.Equal(t, tt.want, got)
		})
	}
}
