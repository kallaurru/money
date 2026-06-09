package money

import (
	i "github.com/kallaurru/money/internal"
	"github.com/shopspring/decimal"
)

type Crypto struct {
	amount   decimal.Decimal
	currency i.CryptoCurrency
}

func NewCrypto(amount decimal.Decimal, code string) (Crypto, bool) {
	// false если не найден был код и вернули описание денег по-умолчанию
	return Crypto{}, true
}
