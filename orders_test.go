package kabustation_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ryomak/kabustation-go"
)

func TestGetOrders_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{
			"ID": "TestID",
			"State": 1,
			"OrderState": 1,
			"OrdType": 1,
			"RecvTime": "2022-01-01T00:00:00Z",
			"Symbol": "TestSymbol",
			"SymbolName": "TestSymbolName",
			"Exchange": 1,
			"ExchangeName": "TestExchange",
			"TimeInForce": 1,
			"Price": 100.0,
			"OrderQty": 100,
			"CumQty": 100,
			"Side": "1",
			"CashMargin": "1",
			"AccountType": 1,
			"DelivType": 1,
			"ExpireDay": 20220101,
			"MarginTradeType": 1,
			"MarginPremium": 100.0,
			"Details": [
				{
					"SeqNum": 1,
					"ID": "TestDetailID",
					"RecType": 1,
					"ExchangeID": "TestExchangeID",
					"State": 1,
					"TransactTime": "2022-01-01T00:00:00Z",
					"OrdType": 1,
					"Price": 100.0,
					"Qty": 100,
					"ExecutionID": "TestExecutionID",
					"ExecutionDay": "2022-01-01T00:00:00Z",
					"DelivDay": 20220101,
					"Commission": 100,
					"CommissionTax": 10
				}
			]
		}`))
	}))
	defer server.Close()

	client := kabustation.NewClient("testPassword")
	client.BaseURL = server.URL

	req := &kabustation.OrdersRequest{
		Product: kabustation.ProductAll,
		ID:      "TestID",
		Uptime:  time.Now(),
		Details: true,
		Symbol:  "TestSymbol",
		State:   kabustation.OrdersRequestStateWait,
	}

	resp, err := client.GetOrders(context.Background(), req)

	if err != nil {
		t.Errorf("GetOrders was incorrect, got: %v", err)
	}
	if resp.ID != "TestID" {
		t.Errorf("GetOrders was incorrect, got: %v, want: %v.", resp.ID, "TestID")
	}
}

func TestGetOrders_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`Error`))
	}))
	defer server.Close()

	client := kabustation.NewClient("testPassword")
	client.BaseURL = server.URL

	req := &kabustation.OrdersRequest{
		Product: kabustation.ProductAll,
		ID:      "TestID",
		Uptime:  time.Now(),
		Details: true,
		Symbol:  "TestSymbol",
		State:   kabustation.OrdersRequestStateWait,
	}

	_, err := client.GetOrders(context.Background(), req)

	if err == nil {
		t.Errorf("GetOrders was incorrect, expected error.")
	}
}

func TestGetOrders_QueryParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		expectedProduct := "0"
		expectedID := "TestID"
		expectedSymbol := "TestSymbol"
		expectedState := "1"

		if req.URL.Query().Get("product") != expectedProduct {
			t.Errorf("Expected product query param %v but got %v", expectedProduct, req.URL.Query().Get("product"))
		}
		if req.URL.Query().Get("id") != expectedID {
			t.Errorf("Expected id query param %v but got %v", expectedID, req.URL.Query().Get("id"))
		}
		if req.URL.Query().Get("symbol") != expectedSymbol {
			t.Errorf("Expected symbol query param %v but got %v", expectedSymbol, req.URL.Query().Get("symbol"))
		}
		if req.URL.Query().Get("state") != expectedState {
			t.Errorf("Expected state query param %v but got %v", expectedState, req.URL.Query().Get("state"))
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{ "ID": "TestID" }`))
	}))
	defer server.Close()

	client := kabustation.NewClient("testPassword")
	client.BaseURL = server.URL

	req := &kabustation.OrdersRequest{
		Product: kabustation.ProductAll,
		ID:      "TestID",
		Uptime:  time.Now(),
		Details: true,
		Symbol:  "TestSymbol",
		State:   kabustation.OrdersRequestStateWait,
	}

	_, err := client.GetOrders(context.Background(), req)

	if err != nil {
		t.Errorf("GetOrders was incorrect, got: %v", err)
	}
}
