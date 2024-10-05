package kabustation

import "fmt"

// MarketCode - 市場コード
type MarketCode int

const (
	MarketCodeTosho   MarketCode = 1  // 東証
	MarketCodeNagoya  MarketCode = 3  // 名証
	MarketCodeFukuoka MarketCode = 5  // 福証
	MarketCodeSapporo MarketCode = 6  // 札証
	MarketCodeDaytime MarketCode = 2  // 日通し
	MarketCodeDay     MarketCode = 23 // 日中
	MarketCodeNight   MarketCode = 24 // 夜間
)

// Symbol - 銘柄
type Symbol string

type SymbolRequest struct {
	Symbol     Symbol
	MarketCode MarketCode
}

func (r *SymbolRequest) symbol() string {
	return fmt.Sprintf("%s@%d", r.Symbol, r.MarketCode)
}

// Product - 商品区分
type Product int

const (
	ProductAll    Product = 0 // すべて
	ProductSpot   Product = 1 // 現物
	ProductCredit Product = 2 // 信用
	ProductFuture Product = 3 // 先物
	ProductOption Product = 4 // OP
)

// Side - 売買区分
type Side string

const (
	SideBuy  Side = "2" // 買い
	SideSell Side = "1" // 売り
)

// CashMargin - 現物/信用区分
type CashMargin string

const (
	CashMarginCash      CashMargin = "1" // 現物
	CashMarginMargin    CashMargin = "2" // 信用
	CashMarginRepayment CashMargin = "3" // 返済
)

// OrderType - 注文種別
type OrderType int

const (
	OrderTypeMarket OrderType = 1 // ザラバ
	OrderTypeOpen   OrderType = 2 // 寄り
	OrderTypeClose  OrderType = 3 // 引け
	OrderTypeUnexec OrderType = 4 // 不成
	OrderTypeLimit  OrderType = 5 // 対当指値
	OrderTypeIOC    OrderType = 6 // IOC
)

type APIResult int

func (a APIResult) IsSuccess() bool {
	return a == 0
}
