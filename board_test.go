package kabustation_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ryomak/kabustation-go"
)

func TestGetBoard_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{
			"Symbol": "TestSymbol",
			"SymbolName": "TestSymbolName",
			"Exchange": 1,
			"ExchangeName": "TestExchange",
			"CurrentPrice": 100,
			"CurrentPriceTime": "2022-01-01T00:00:00Z",
			"CurrentPriceChangeStatus": "TestStatus",
			"CurrentPriceStatus": 1,
			"CalcPrice": 100.0,
			"PreviousClose": 100,
			"PreviousCloseTime": "2022-01-01T00:00:00Z",
			"ChangePreviousClose": 0,
			"ChangePreviousClosePer": 0.0,
			"OpeningPrice": 100,
			"OpeningPriceTime": "2022-01-01T00:00:00Z",
			"HighPrice": 100,
			"HighPriceTime": "2022-01-01T00:00:00Z",
			"LowPrice": 100,
			"LowPriceTime": "2022-01-01T00:00:00Z",
			"TradingVolume": 100,
			"TradingVolumeTime": "2022-01-01T00:00:00Z",
			"VWAP": 100.0,
			"TradingValue": 10000,
			"BidQty": 100,
			"BidPrice": 100.0,
			"BidTime": "2022-01-01T00:00:00Z",
			"BidSign": "TestSign",
			"MarketOrderSellQty": 100,
			"Sell1": {"Time": "2022-01-01T00:00:00Z", "Sign": "TestSign", "Price": 100.0, "Qty": 100},
			"Sell2": {"Price": 100, "Qty": 100},
			"Sell3": {"Price": 100.0, "Qty": 100},
			"Sell4": {"Price": 100, "Qty": 100},
			"Sell5": {"Price": 100.0, "Qty": 100},
			"Sell6": {"Price": 100, "Qty": 100},
			"Sell7": {"Price": 100.0, "Qty": 100},
			"Sell8": {"Price": 100, "Qty": 100},
			"Sell9": {"Price": 100.0, "Qty": 100},
			"Sell10": {"Price": 100, "Qty": 100},
			"AskQty": 100,
			"AskPrice": 100.0,
			"AskTime": "2022-01-01T00:00:00Z",
			"AskSign": "TestSign",
			"MarketOrderBuyQty": 100,
			"Buy1": {"Time": "2022-01-01T00:00:00Z", "Sign": "TestSign", "Price": 100.0, "Qty": 100},
			"Buy2": {"Price": 100, "Qty": 100},
			"Buy3": {"Price": 100.0, "Qty": 100},
			"Buy4": {"Price": 100, "Qty": 100},
			"Buy5": {"Price": 100.0, "Qty": 100},
			"Buy6": {"Price": 100, "Qty": 100},
			"Buy7": {"Price": 100.0, "Qty": 100},
			"Buy8": {"Price": 100, "Qty": 100},
			"Buy9": {"Price": 100.0, "Qty": 100},
			"Buy10": {"Price": 100, "Qty": 100},
			"OverSellQty": 100,
			"UnderBuyQty": 100,
			"TotalMarketValue": 1000000.0,
			"SecurityType": 1
		}`))
	}))
	defer server.Close()

	client := kabustation.NewClient("testPassword")
	client.BaseURL = server.URL

	req := &kabustation.SymbolRequest{
		Symbol: "TestSymbol",
	}

	resp, err := client.GetBoard(context.Background(), req)

	if err != nil {
		t.Errorf("GetBoard was incorrect, got: %v", err)
	}
	if resp.Symbol != "TestSymbol" {
		t.Errorf("GetBoard was incorrect, got: %v, want: %v.", resp.Symbol, "TestSymbol")
	}
}

func TestGetBoard_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`Error`))
	}))
	defer server.Close()

	client := kabustation.NewClient("testPassword")
	client.BaseURL = server.URL

	req := &kabustation.SymbolRequest{
		Symbol: "TestSymbol",
	}

	_, err := client.GetBoard(context.Background(), req)

	if err == nil {
		t.Errorf("GetBoard was incorrect, expected error.")
	}
}
