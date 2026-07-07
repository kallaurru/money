package money

import (
	cr "github.com/kallaurru/money/internal/crypto"
	crn "github.com/kallaurru/money/internal/currency"
	"github.com/shopspring/decimal"
)

type Money struct {
	amount   decimal.Decimal
	currency Monetary
}

func NewMoney(amount decimal.Decimal, code string) (Money, bool) {
	val, ok := crn.FindByCode(code)
	if !ok {
		return Money{}, ok
	}

	return Money{
		amount:   amount,
		currency: val,
	}, true
}

func NewCrypto(amount decimal.Decimal, code string) (Money, bool) {
	val, ok := cr.FindByCode(code)
	if !ok {
		return Money{}, ok
	}
	return Money{
		amount:   amount,
		currency: val,
	}, true
}

func (m Money) Format() string {
	return m.currency.Format(m.amount)
}

func (m Money) IsFiat() bool {
	return m.currency.IsFiat()
}
