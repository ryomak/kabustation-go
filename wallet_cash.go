package kabustation

import (
	"context"
	"net/http"
)

type WalletCashResponse struct {
	StockAccountWallet      *float64 `json:"StockAccountWallet"`
	AuKCStockAccountWallet  *float64 `json:"AuKCStockAccountWallet"`
	AuJbnStockAccountWallet *float64 `json:"AuJbnStockAccountWallet"`
}

func (c *Client) WalletCash(ctx context.Context) (*WalletCashResponse, error) {
	req, err := c.newRequest(http.MethodGet, "/wallet/cash")
	if err != nil {
		return nil, err
	}
	var resp WalletCashResponse
	if _, err := c.do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type WalletCashSymbolResponse struct {
	StockAccountWallet      *float64 `json:"StockAccountWallet"`
	AuKCStockAccountWallet  *float64 `json:"AuKCStockAccountWallet"`
	AuJbnStockAccountWallet *float64 `json:"AuJbnStockAccountWallet"`
}

func (c *Client) WalletCashSymbol(ctx context.Context, req *SymbolRequest) (*WalletCashSymbolResponse, error) {
	request, err := c.newRequest(http.MethodGet, "/wallet/cash/"+req.symbol())
	if err != nil {
		return nil, err
	}
	var resp WalletCashSymbolResponse
	if _, err := c.do(withRateLimiterContext(ctx, rateLimitPathWallet), request, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
