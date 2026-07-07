package money

import "github.com/shopspring/decimal"

type Monetary interface {
	Format(decimal.Decimal) string
	IsFiat() bool
}
