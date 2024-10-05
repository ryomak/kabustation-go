package kabustation_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ryomak/kabustation-go"
)

func TestWalletMargin_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{
			"MarginAccountWallet": 100.0,
			"DepositkeepRate": 0.5,
			"ConsignmentDepositRate": 0.5,
			"CashOfConsignmentDepositRate": 0.5
		}`))
	}))
	defer server.Close()

	client := kabustation.NewClient(
		"testPassword",
		kabustation.WithBaseURL(server.URL),
	)

	resp, err := client.WalletMargin(context.Background())

	if err != nil {
		t.Errorf("WalletMargin was incorrect, got: %v", err)
	}
	if *resp.MarginAccountWallet != 100.0 {
		t.Errorf("WalletMargin was incorrect, got: %v, want: %v.", *resp.MarginAccountWallet, 100.0)
	}
}

func TestWalletMargin_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`Error`))
	}))
	defer server.Close()

	client := kabustation.NewClient(
		"testPassword",
		kabustation.WithBaseURL(server.URL),
	)

	_, err := client.WalletMargin(context.Background())

	if err == nil {
		t.Errorf("WalletMargin was incorrect, expected error.")
	}
}

func TestWalletMarginSymbol_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{
			"MarginAccountWallet": 100.0,
			"DepositkeepRate": 0.5,
			"ConsignmentDepositRate": 0.5,
			"CashOfConsignmentDepositRate": 0.5
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

	resp, err := client.WalletMarginSymbol(context.Background(), req)

	if err != nil {
		t.Errorf("WalletMarginSymbol was incorrect, got: %v", err)
	}
	if *resp.MarginAccountWallet != 100.0 {
		t.Errorf("WalletMarginSymbol was incorrect, got: %v, want: %v.", *resp.MarginAccountWallet, 100.0)
	}
}

func TestWalletMarginSymbol_Error(t *testing.T) {
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

	_, err := client.WalletMarginSymbol(context.Background(), req)

	if err == nil {
		t.Errorf("WalletMarginSymbol was incorrect, expected error.")
	}
}
