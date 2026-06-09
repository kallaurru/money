package money

import "github.com/shopspring/decimal"

// заготовка под очень мелкие значения. Например цена акций 0.00374 или цена крипты к доллару 43,9873

func NewPenny(val int64, std bool) decimal.Decimal {
	if std {
		return decimal.New(val, 6)
	}
	return decimal.New(val, 8)
}
