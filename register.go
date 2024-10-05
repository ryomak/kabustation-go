package kabustation

import (
	"context"
	"net/http"
)

type RegisterRequest struct {
	Symbols []RegisterRequestSymbol `json:"Symbols"`
}

type RegisterRequestSymbol struct {
	Symbol   string     `json:"Symbol"`
	Exchange MarketCode `json:"Exchange"`
}

type RegisterResponse struct {
	RegistList []struct {
		Symbol   string     `json:"Symbol"`
		Exchange MarketCode `json:"Exchange"`
	} `json:"RegistList"`
}

func (c *Client) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	request, err := c.newRequest(http.MethodPost, "/register", withBody(req))
	if err != nil {
		return nil, err
	}
	var resp RegisterResponse
	if _, err := c.do(withRateLimiterContext(ctx, rateLimitPathRegister), request, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) Unregister(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	request, err := c.newRequest(http.MethodPost, "/unregister", withBody(req))
	if err != nil {
		return nil, err
	}
	var resp RegisterResponse
	if _, err := c.do(withRateLimiterContext(ctx, rateLimitPathRegister), request, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
