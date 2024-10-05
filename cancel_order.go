package kabustation

import (
	"context"
	"net/http"
)

type CancelOrderRequest struct {
	OrderId  string `json:"OrderId"`
	Password string `json:"Password"`
}

type CancelOrderResponse struct {
	Result  APIResult `json:"Result"`
	OrderId string    `json:"OrderId"`
}

func (c *CancelOrderResponse) IsSuccess() bool {
	return c.Result == 0
}

func (c *Client) CancelOrder(ctx context.Context, req *CancelOrderRequest) (*CancelOrderResponse, error) {
	request, err := c.newRequest(http.MethodPut, "/cancelorder", withBody(req))
	if err != nil {
		return nil, err
	}
	var resp CancelOrderResponse
	if _, err := c.do(withRateLimiterContext(ctx, rateLimitPathOrder), request, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
