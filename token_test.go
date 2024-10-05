package kabustation_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ryomak/kabustation-go"
)

func TestGetToken_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{
			"ResultCode": 0,
			"Token": "testToken"
		}`))
	}))
	defer server.Close()

	client := kabustation.NewClient(
		"testPassword",
		kabustation.WithBaseURL(server.URL),
	)

	resp, err := client.GetToken(context.Background(), false, false)

	if err != nil {
		t.Errorf("GetToken was incorrect, got: %v", err)
	}
	if resp.Token != "testToken" {
		t.Errorf("GetToken was incorrect, got: %v, want: %v.", resp.Token, "testToken")
	}
}

func TestGetToken_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`Error`))
	}))
	defer server.Close()

	client := kabustation.NewClient(
		"testPassword",
		kabustation.WithBaseURL(server.URL),
	)

	_, err := client.GetToken(context.Background(), false, false)

	if err == nil {
		t.Errorf("GetToken was incorrect, expected error.")
	}
}

func TestGetToken_Cache(t *testing.T) {
	var callCount int
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		callCount++
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{
   "ResultCode": 0,
   "Token": "testToken"
  }`))
	}))
	defer server.Close()

	client := kabustation.NewClient(
		"testPassword",
		kabustation.WithBaseURL(server.URL),
	)

	// First call to GetToken, should hit the API
	_, _ = client.GetToken(context.Background(), false, true)
	if callCount != 1 {
		t.Errorf("GetToken was incorrect, expected API to be hit once, got: %v", callCount)
	}

	// Second call to GetToken, should use the cache and not hit the API
	_, _ = client.GetToken(context.Background(), false, true)
	if callCount != 1 {
		t.Errorf("GetToken was incorrect, expected API to be hit once, got: %v", callCount)
	}

	// Third call to GetToken with force, should hit the API
	_, _ = client.GetToken(context.Background(), true, true)
	if callCount != 2 {
		t.Errorf("GetToken was incorrect, expected API to be hit twice, got: %v", callCount)
	}
}
