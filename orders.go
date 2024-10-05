package kabustation

import (
	"context"
	"net/http"
	"time"
)

type OrdersRequest struct {
	Product Product     `url:"product"`
	ID      string      `url:"id"`
	Uptime  time.Time   `url:"uptime"`
	Details bool        `url:"details"`
	Symbol  string      `url:"symbol"`
	State   OrdersState `url:"state"`
}

// OrdersState - 注文状態
type OrdersState int

const (
	OrdersRequestStateWait                    OrdersState = 1 // 待機（発注待機）
	OrdersRequestStateProcessing              OrdersState = 2 // 処理中（発注送信中）
	OrdersRequestStateProcessed               OrdersState = 3 // 処理済（発注済・訂正済）
	OrdersRequestStateCorrectionCancelSending OrdersState = 4 // 訂正取消送信中
)

type OrdersResponse struct {
	ID              string      `json:"ID"`
	State           OrdersState `json:"State"`
	OrderState      OrdersState `json:"OrderState"`
	OrdType         OrderType   `json:"OrdType"`
	RecvTime        time.Time   `json:"RecvTime"`
	Symbol          string      `json:"Symbol"`
	SymbolName      string      `json:"SymbolName"`
	Exchange        int         `json:"Exchange"`
	ExchangeName    string      `json:"ExchangeName"`
	TimeInForce     int         `json:"TimeInForce"`
	Price           float64     `json:"Price"`
	OrderQty        int         `json:"OrderQty"`
	CumQty          int         `json:"CumQty"`
	Side            Side        `json:"Side"`
	CashMargin      CashMargin  `json:"CashMargin"`
	AccountType     int         `json:"AccountType"`
	DelivType       int         `json:"DelivType"`
	ExpireDay       int         `json:"ExpireDay"`
	MarginTradeType int         `json:"MarginTradeType"`
	MarginPremium   interface{} `json:"MarginPremium"`
	Details         []struct {
		SeqNum        int       `json:"SeqNum"`
		ID            string    `json:"ID"`
		RecType       int       `json:"RecType"`
		ExchangeID    string    `json:"ExchangeID"`
		State         int       `json:"State"`
		TransactTime  time.Time `json:"TransactTime"`
		OrdType       OrderType `json:"OrdType"`
		Price         float64   `json:"Price"`
		Qty           int       `json:"Qty"`
		ExecutionID   string    `json:"ExecutionID"`
		ExecutionDay  time.Time `json:"ExecutionDay"`
		DelivDay      int       `json:"DelivDay"`
		Commission    int       `json:"Commission"`
		CommissionTax int       `json:"CommissionTax"`
	} `json:"Details"`
}

func (c *Client) GetOrders(ctx context.Context, req *OrdersRequest) (*OrdersResponse, error) {

	request, err := c.newRequest(http.MethodGet, "/orders", withQueryParams(req))
	if err != nil {
		return nil, err
	}
	var resp OrdersResponse
	if _, err := c.do(withRateLimiterContext(ctx, rateLimitPathInfo), request, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
