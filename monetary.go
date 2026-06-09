package money

import "github.com/shopspring/decimal"

type Monetary interface {
	Amount() decimal.Decimal
	String() string
	FmtAmount() string
}
