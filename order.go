package gofastmag

import (
	"github.com/shopspring/decimal"
	"time"
)

type OrderLine struct {
	StoreCode    string
	Color        string
	BarCode      string
	Size         string
	SaleId       int
	StockId      int
	Origin       string
	ReceiptAt    time.Time
	SeasonOrigin string
	Designation  string
	Designation2 string
	Supplier     string
	Family       string
	Season       string
	BuyingPrice  decimal.Decimal
	SellingPrice decimal.Decimal
	LineId       int
	Price        decimal.Decimal
	Quantity     decimal.Decimal
	Discount     float64
	Total        decimal.Decimal
	Reason string
	VAT float64
	Remainder decimal.Decimal
	OrderId int
}