package kabustation_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ryomak/kabustation-go"
)

func TestWalletCash_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{
   "StockAccountWallet": 100.0,
   "AuKCStockAccountWallet": 50.0,
   "AuJbnStockAccountWallet": 25.0
  }`))
	}))
	defer server.Close()

	client := kabustation.NewClient(
		"testPassword",
		kabustation.WithBaseURL(server.URL),
	)

	resp, err := client.WalletCash(context.Background())

	if err != nil {
		t.Errorf("WalletCash was incorrect, got: %v", err)
	}
	if *resp.StockAccountWallet != 100.0 {
		t.Errorf("WalletCash was incorrect, got: %v, want: %v.", *resp.StockAccountWallet, 100.0)
	}
}

func TestWalletCash_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`Error`))
	}))
	defer server.Close()

	client := kabustation.NewClient(
		"testPassword",
		kabustation.WithBaseURL(server.URL),
	)

	_, err := client.WalletCash(context.Background())

	if err == nil {
		t.Errorf("WalletCash was incorrect, expected error.")
	}
}

func TestWalletCashSymbol_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{
   "StockAccountWallet": 100.0,
   "AuKCStockAccountWallet": 50.0,
   "AuJbnStockAccountWallet": 25.0
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

	resp, err := client.WalletCashSymbol(context.Background(), req)

	if err != nil {
		t.Errorf("WalletCashSymbol was incorrect, got: %v", err)
	}
	if *resp.StockAccountWallet != 100.0 {
		t.Errorf("WalletCashSymbol was incorrect, got: %v, want: %v.", *resp.StockAccountWallet, 100.0)
	}
}

func TestWalletCashSymbol_Error(t *testing.T) {
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

	_, err := client.WalletCashSymbol(context.Background(), req)

	if err == nil {
		t.Errorf("WalletCashSymbol was incorrect, expected error.")
	}
}
