package money

import (
	i "github.com/kallaurru/money/internal"
	"github.com/shopspring/decimal"
)

type Money struct {
	amount   decimal.Decimal
	currency i.Currency
}

func NewMoney(amount decimal.Decimal, code string) (Money, bool) {
	// false если не найден был код и вернули описание денег по-умолчанию
	return Money{}, true
}
