package gofastmag

type LineType string

const (
	SaleHeaderLineType LineType = "ENTETE"
	SaleLineLineType            = "LIGNE"
)

type TransactionType string

const (
	SaleTransactionType TransactionType = "VENTE"
)

type StateOfPiece string

const (
	L StateOfPiece = "L"
	P              = "P"
	E              = "E"
)
