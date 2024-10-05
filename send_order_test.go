package kabustation_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ryomak/kabustation-go"
)

func TestSendOrder_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{
			"Result": 0,
			"OrderId": "testOrderId"
		}`))
	}))
	defer server.Close()

	client := kabustation.NewClient("testPassword")
	client.BaseURL = server.URL

	req := &kabustation.SendOrderRequest{
		Password: "testPassword",
		Symbol:   "TestSymbol",
		Exchange: 1,
	}

	resp, err := client.SendOrder(context.Background(), req)

	if err != nil {
		t.Errorf("SendOrder was incorrect, got: %v", err)
	}
	if resp.OrderId != "testOrderId" {
		t.Errorf("SendOrder was incorrect, got: %v, want: %v.", resp.OrderId, "testOrderId")
	}
}

func TestSendOrder_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`Error`))
	}))
	defer server.Close()

	client := kabustation.NewClient("testPassword")
	client.BaseURL = server.URL

	req := &kabustation.SendOrderRequest{
		Password: "testPassword",
		Symbol:   "TestSymbol",
		Exchange: 1,
	}

	_, err := client.SendOrder(context.Background(), req)

	if err == nil {
		t.Errorf("SendOrder was incorrect, expected error.")
	}
}
