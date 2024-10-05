package kabustation

import (
	"context"
	"net/http"
)

type SendOrderRequest struct {
	Password          string     `json:"Password"`
	Symbol            string     `json:"Symbol"`
	Exchange          int        `json:"Exchange"`
	SecurityType      int        `json:"SecurityType"`
	Side              Side       `json:"Side"`
	CashMargin        CashMargin `json:"CashMargin"`
	MarginTradeType   int        `json:"MarginTradeType"`
	MarginPremiumUnit float64    `json:"MarginPremiumUnit"`
	DelivType         int        `json:"DelivType"`
	AccountType       int        `json:"AccountType"`
	Qty               int        `json:"Qty"`
	ClosePositions    []struct {
		HoldID string `json:"HoldID"`
		Qty    int    `json:"Qty"`
	} `json:"ClosePositions"`
	FrontOrderType    int `json:"FrontOrderType"`
	ExpireDay         int `json:"ExpireDay"`
	ReverseLimitOrder struct {
		TriggerSec        int `json:"TriggerSec"`
		TriggerPrice      int `json:"TriggerPrice"`
		UnderOver         int `json:"UnderOver"`
		AfterHitOrderType int `json:"AfterHitOrderType"`
		AfterHitPrice     int `json:"AfterHitPrice"`
	} `json:"ReverseLimitOrder"`
}

type SendOrderResponse struct {
	Result  int    `json:"Result"`
	OrderId string `json:"OrderId"`
}

func (s *SendOrderResponse) IsSuccess() bool {
	return s.Result == 0
}

func (c *Client) SendOrder(ctx context.Context, req *SendOrderRequest) (*SendOrderResponse, error) {
	request, err := c.newRequest(http.MethodPost, "/sendorder", withBody(req))
	if err != nil {
		return nil, err
	}
	var resp SendOrderResponse
	if _, err := c.do(withRateLimiterContext(ctx, rateLimitPathOrder), request, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
