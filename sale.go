package gofastmag

import "github.com/shopspring/decimal"

type Sale struct {
	Store               string
	TotalAmount         decimal.Decimal
	TotalQuantity       int
	TransactionType     TransactionType
	Seller              string
	Email               string
	ShippingAddressName string
	User                string
	PieceStatus         string
	IsWithoutTaxes      bool
	PieceState          string
	Comment             string
	SaleFrom            string
	PickupPoint         string
	PaymentCode         string
	Event               string
}

type SaleLine struct {
	ProductRef         string
	Size               string
	Color              string
	Designation        string
	UnitPriceWithTaxes decimal.Decimal
	Quantity           int
	Discount           decimal.Decimal
	LinePriceWithTaxes decimal.Decimal
	Reason             string
	Comment            string
	OrderLine          decimal.Decimal
	Designation2       string
	Pme                string
}
