package lotusgo

import (
	"net/http"
	"reflect"
	"testing"
)

func TestClient_ListCustomers(t *testing.T) {
	type fields struct {
		BaseURL    string
		apiKey     string
		HTTPClient *http.Client
		debug      bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    []ListCustomerResponse
		wantErr bool
	}{
		// TODO: Add test cases.
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
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ListCustomers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.ListCustomers() = %v, want %v", got, tt.want)
			}
		})
	}
}
