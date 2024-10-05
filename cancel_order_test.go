package kabustation_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ryomak/kabustation-go"
)

func TestCancelOrder_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{
			"Result": 0,
			"OrderId": "testOrderId"
		}`))
	}))
	defer server.Close()

	client := kabustation.NewClient(
		"testPassword",
		kabustation.WithBaseURL(server.URL),
	)

	req := &kabustation.CancelOrderRequest{
		OrderId:  "testOrderId",
		Password: "testPassword",
	}

	resp, err := client.CancelOrder(context.Background(), req)

	if err != nil {
		t.Errorf("CancelOrder was incorrect, got: %v", err)
	}
	if resp.OrderId != "testOrderId" {
		t.Errorf("CancelOrder was incorrect, got: %v, want: %v.", resp.OrderId, "testOrderId")
	}
}

func TestCancelOrder_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`Error`))
	}))
	defer server.Close()

	client := kabustation.NewClient(
		"testPassword",
		kabustation.WithBaseURL(server.URL),
	)

	req := &kabustation.CancelOrderRequest{
		OrderId:  "testOrderId",
		Password: "testPassword",
	}

	_, err := client.CancelOrder(context.Background(), req)

	if err == nil {
		t.Errorf("CancelOrder was incorrect, expected error.")
	}
}
