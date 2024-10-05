package kabustation_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ryomak/kabustation-go"
)

func TestRegister_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{
			"RegistList": [
				{
					"Symbol": "TestSymbol",
					"Exchange": 1
				}
			]
		}`))
	}))
	defer server.Close()

	client := kabustation.NewClient("testPassword")
	client.BaseURL = server.URL

	req := &kabustation.RegisterRequest{
		Symbols: []kabustation.RegisterRequestSymbol{
			{
				Symbol:   "TestSymbol",
				Exchange: kabustation.MarketCodeTosho,
			},
		},
	}

	resp, err := client.Register(context.Background(), req)

	if err != nil {
		t.Errorf("Register was incorrect, got: %v", err)
	}
	if resp.RegistList[0].Symbol != "TestSymbol" {
		t.Errorf("Register was incorrect, got: %v, want: %v.", resp.RegistList[0].Symbol, "TestSymbol")
	}
}

func TestRegister_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`Error`))
	}))
	defer server.Close()

	client := kabustation.NewClient("testPassword")
	client.BaseURL = server.URL

	req := &kabustation.RegisterRequest{
		Symbols: []kabustation.RegisterRequestSymbol{
			{
				Symbol:   "TestSymbol",
				Exchange: kabustation.MarketCodeTosho,
			},
		},
	}

	_, err := client.Register(context.Background(), req)

	if err == nil {
		t.Errorf("Register was incorrect, expected error.")
	}
}

func TestUnregister_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{
			"RegistList": []
		}`))
	}))
	defer server.Close()

	client := kabustation.NewClient("testPassword")
	client.BaseURL = server.URL

	req := &kabustation.RegisterRequest{
		Symbols: []kabustation.RegisterRequestSymbol{
			{
				Symbol:   "TestSymbol",
				Exchange: kabustation.MarketCodeTosho,
			},
		},
	}

	resp, err := client.Unregister(context.Background(), req)

	if err != nil {
		t.Errorf("Unregister was incorrect, got: %v", err)
	}
	if len(resp.RegistList) != 0 {
		t.Errorf("Unregister was incorrect, got: %v, want: %v.", len(resp.RegistList), 0)
	}
}

func TestUnregister_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`Error`))
	}))
	defer server.Close()

	client := kabustation.NewClient("testPassword")
	client.BaseURL = server.URL

	req := &kabustation.RegisterRequest{
		Symbols: []kabustation.RegisterRequestSymbol{
			{
				Symbol:   "TestSymbol",
				Exchange: kabustation.MarketCodeTosho,
			},
		},
	}

	_, err := client.Unregister(context.Background(), req)

	if err == nil {
		t.Errorf("Unregister was incorrect, expected error.")
	}
}
