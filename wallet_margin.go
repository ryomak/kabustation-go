package kabustation

import (
	"context"
	"net/http"
)

type WalletMarginResponse struct {
	MarginAccountWallet          *float64 `json:"MarginAccountWallet"`
	DepositkeepRate              *float64 `json:"DepositkeepRate"`
	ConsignmentDepositRate       *float64 `json:"ConsignmentDepositRate"`
	CashOfConsignmentDepositRate *float64 `json:"CashOfConsignmentDepositRate"`
}

func (c *Client) WalletMargin(ctx context.Context) (*WalletMarginResponse, error) {
	req, err := c.newRequest(http.MethodGet, "/wallet/margin")
	if err != nil {
		return nil, err
	}
	var resp WalletMarginResponse
	if _, err := c.do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type WalletMarginSymbolResponse struct {
	MarginAccountWallet          *float64 `json:"MarginAccountWallet"`
	DepositkeepRate              *float64 `json:"DepositkeepRate"`
	ConsignmentDepositRate       *float64 `json:"ConsignmentDepositRate"`
	CashOfConsignmentDepositRate *float64 `json:"CashOfConsignmentDepositRate"`
}

func (c *Client) WalletMarginSymbol(ctx context.Context, req *SymbolRequest) (*WalletMarginSymbolResponse, error) {
	request, err := c.newRequest(http.MethodGet, "/wallet/margin/"+req.symbol())
	if err != nil {
		return nil, err
	}
	var resp WalletMarginSymbolResponse
	if _, err := c.do(withRateLimiterContext(ctx, rateLimitPathWallet), request, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
