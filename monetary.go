package money

import "github.com/shopspring/decimal"

type Monetary interface {
	IsFiat() bool
	Amount() decimal.Decimal
}

type CoinDescriber interface {
	Format(decimal.Decimal) string
	String(decimal.Decimal) string
	Advanced() string
}
