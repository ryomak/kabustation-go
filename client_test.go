package kabustation_test

import (
	"net/http"
	"testing"

	"github.com/ryomak/kabustation-go"
)

func TestNewClient(t *testing.T) {
	client := kabustation.NewClient("testPassword")
	if client == nil {
		t.Errorf("NewClient was incorrect, got: nil")
	}
}

func TestWithHTTPClient(t *testing.T) {
	httpClient := &http.Client{}
	option := kabustation.WithHTTPClient(httpClient)
	client := kabustation.NewClient("testPassword", option)
	if client.HTTPClient != httpClient {
		t.Errorf("WithHTTPClient was incorrect, got: %v, want: %v.", client.HTTPClient, httpClient)
	}
}

func TestWithRateLimiter(t *testing.T) {
	option := kabustation.WithRateLimiter()
	client := kabustation.NewClient("testPassword", option)
	if client.HTTPClient.Transport == nil {
		t.Errorf("WithRateLimiter was incorrect, got: nil")
	}
}
