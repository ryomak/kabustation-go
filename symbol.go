package kabustation

import (
	"context"
	"net/http"
)

type SymbolResponse struct {
	Symbol          string  `json:"Symbol"`
	SymbolName      string  `json:"SymbolName"`
	DisplayName     string  `json:"DisplayName"`
	Exchange        int     `json:"Exchange"`
	ExchangeName    string  `json:"ExchangeName"`
	TradingUnit     int     `json:"TradingUnit"`
	PriceRangeGroup string  `json:"PriceRangeGroup"`
	UpperLimit      int     `json:"UpperLimit"`
	LowerLimit      int     `json:"LowerLimit"`
	Underlyer       string  `json:"Underlyer"`
	DerivMonth      string  `json:"DerivMonth"`
	TradeEnd        int     `json:"TradeEnd"`
	TradeStart      int     `json:"TradeStart"`
	ClearingPrice   float64 `json:"ClearingPrice"`
}

func (c *Client) GetSymbol(ctx context.Context, req *SymbolRequest) (*SymbolResponse, error) {
	request, err := c.newRequest(http.MethodGet, "/symbol/"+req.symbol(), withQueryParams(req))
	if err != nil {
		return nil, err
	}
	var resp SymbolResponse
	if _, err := c.do(withRateLimiterContext(ctx, rateLimitPathInfo), request, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
