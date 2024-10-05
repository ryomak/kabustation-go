package kabustation_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ryomak/kabustation-go"
)

func TestGetSymbol_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{
   "Symbol": "TestSymbol",
   "SymbolName": "TestSymbolName",
   "DisplayName": "TestDisplayName",
   "Exchange": 1,
   "ExchangeName": "TestExchange",
   "TradingUnit": 100,
   "PriceRangeGroup": "TestPriceRangeGroup",
   "UpperLimit": 1000,
   "LowerLimit": 100,
   "Underlyer": "TestUnderlyer",
   "DerivMonth": "2022-01",
   "TradeEnd": 1500,
   "TradeStart": 900,
   "ClearingPrice": 500.0
  }`))
	}))
	defer server.Close()

	client := kabustation.NewClient(
		"testPassword",
		kabustation.WithBaseURL(server.URL),
	)

	req := &kabustation.SymbolRequest{
		Symbol: "TestSymbol",
	}

	resp, err := client.GetSymbol(context.Background(), req)

	if err != nil {
		t.Errorf("GetSymbol was incorrect, got: %v", err)
	}
	if resp.Symbol != "TestSymbol" {
		t.Errorf("GetSymbol was incorrect, got: %v, want: %v.", resp.Symbol, "TestSymbol")
	}
}

func TestGetSymbol_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`Error`))
	}))
	defer server.Close()

	client := kabustation.NewClient(
		"testPassword",
		kabustation.WithBaseURL(server.URL),
	)

	req := &kabustation.SymbolRequest{
		Symbol: "TestSymbol",
	}

	_, err := client.GetSymbol(context.Background(), req)

	if err == nil {
		t.Errorf("GetSymbol was incorrect, expected error.")
	}
}

func TestGetSymbol_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := kabustation.NewClient(
		"testPassword",
		kabustation.WithBaseURL(server.URL),
	)

	req := &kabustation.SymbolRequest{
		Symbol: "InvalidSymbol",
	}

	_, err := client.GetSymbol(context.Background(), req)

	if err == nil {
		t.Errorf("GetSymbol was incorrect, expected error.")
	}
}
