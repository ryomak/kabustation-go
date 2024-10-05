package kabustation

import (
	"context"
	"net/http"
	"time"
)

type BoardResponse struct {
	Symbol                   string    `json:"Symbol"`
	SymbolName               string    `json:"SymbolName"`
	Exchange                 int       `json:"Exchange"`
	ExchangeName             string    `json:"ExchangeName"`
	CurrentPrice             int       `json:"CurrentPrice"`
	CurrentPriceTime         time.Time `json:"CurrentPriceTime"`
	CurrentPriceChangeStatus string    `json:"CurrentPriceChangeStatus"`
	CurrentPriceStatus       int       `json:"CurrentPriceStatus"`
	CalcPrice                float64   `json:"CalcPrice"`
	PreviousClose            int       `json:"PreviousClose"`
	PreviousCloseTime        time.Time `json:"PreviousCloseTime"`
	ChangePreviousClose      int       `json:"ChangePreviousClose"`
	ChangePreviousClosePer   float64   `json:"ChangePreviousClosePer"`
	OpeningPrice             int       `json:"OpeningPrice"`
	OpeningPriceTime         time.Time `json:"OpeningPriceTime"`
	HighPrice                int       `json:"HighPrice"`
	HighPriceTime            time.Time `json:"HighPriceTime"`
	LowPrice                 int       `json:"LowPrice"`
	LowPriceTime             time.Time `json:"LowPriceTime"`
	TradingVolume            int       `json:"TradingVolume"`
	TradingVolumeTime        time.Time `json:"TradingVolumeTime"`
	VWAP                     float64   `json:"VWAP"`
	TradingValue             int64     `json:"TradingValue"`
	BidQty                   int       `json:"BidQty"`
	BidPrice                 float64   `json:"BidPrice"`
	BidTime                  time.Time `json:"BidTime"`
	BidSign                  string    `json:"BidSign"`
	MarketOrderSellQty       int       `json:"MarketOrderSellQty"`
	Sell1                    struct {
		Time  time.Time `json:"Time"`
		Sign  string    `json:"Sign"`
		Price float64   `json:"Price"`
		Qty   int       `json:"Qty"`
	} `json:"Sell1"`
	Sell2 struct {
		Price int `json:"Price"`
		Qty   int `json:"Qty"`
	} `json:"Sell2"`
	Sell3 struct {
		Price float64 `json:"Price"`
		Qty   int     `json:"Qty"`
	} `json:"Sell3"`
	Sell4 struct {
		Price int `json:"Price"`
		Qty   int `json:"Qty"`
	} `json:"Sell4"`
	Sell5 struct {
		Price float64 `json:"Price"`
		Qty   int     `json:"Qty"`
	} `json:"Sell5"`
	Sell6 struct {
		Price int `json:"Price"`
		Qty   int `json:"Qty"`
	} `json:"Sell6"`
	Sell7 struct {
		Price float64 `json:"Price"`
		Qty   int     `json:"Qty"`
	} `json:"Sell7"`
	Sell8 struct {
		Price int `json:"Price"`
		Qty   int `json:"Qty"`
	} `json:"Sell8"`
	Sell9 struct {
		Price float64 `json:"Price"`
		Qty   int     `json:"Qty"`
	} `json:"Sell9"`
	Sell10 struct {
		Price int `json:"Price"`
		Qty   int `json:"Qty"`
	} `json:"Sell10"`
	AskQty            int       `json:"AskQty"`
	AskPrice          float64   `json:"AskPrice"`
	AskTime           time.Time `json:"AskTime"`
	AskSign           string    `json:"AskSign"`
	MarketOrderBuyQty int       `json:"MarketOrderBuyQty"`
	Buy1              struct {
		Time  time.Time `json:"Time"`
		Sign  string    `json:"Sign"`
		Price float64   `json:"Price"`
		Qty   int       `json:"Qty"`
	} `json:"Buy1"`
	Buy2 struct {
		Price int `json:"Price"`
		Qty   int `json:"Qty"`
	} `json:"Buy2"`
	Buy3 struct {
		Price float64 `json:"Price"`
		Qty   int     `json:"Qty"`
	} `json:"Buy3"`
	Buy4 struct {
		Price int `json:"Price"`
		Qty   int `json:"Qty"`
	} `json:"Buy4"`
	Buy5 struct {
		Price float64 `json:"Price"`
		Qty   int     `json:"Qty"`
	} `json:"Buy5"`
	Buy6 struct {
		Price int `json:"Price"`
		Qty   int `json:"Qty"`
	} `json:"Buy6"`
	Buy7 struct {
		Price float64 `json:"Price"`
		Qty   int     `json:"Qty"`
	} `json:"Buy7"`
	Buy8 struct {
		Price int `json:"Price"`
		Qty   int `json:"Qty"`
	} `json:"Buy8"`
	Buy9 struct {
		Price float64 `json:"Price"`
		Qty   int     `json:"Qty"`
	} `json:"Buy9"`
	Buy10 struct {
		Price int `json:"Price"`
		Qty   int `json:"Qty"`
	} `json:"Buy10"`
	OverSellQty      int     `json:"OverSellQty"`
	UnderBuyQty      int     `json:"UnderBuyQty"`
	TotalMarketValue float64 `json:"TotalMarketValue"`
	SecurityType     int     `json:"SecurityType"`
}

func (c *Client) GetBoard(ctx context.Context, req *SymbolRequest) (*BoardResponse, error) {
	request, err := c.newRequest(http.MethodGet, "/board/"+req.symbol())
	if err != nil {
		return nil, err
	}
	var resp BoardResponse
	if _, err := c.do(withRateLimiterContext(ctx, rateLimitPathInfo), request, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
