package crypto

import (
	// gm "github.com/Rhymond/go-money"
	f "github.com/kallaurru/money/internal/formatter"
	"github.com/shopspring/decimal"
)

var cc = Cryptos{
	BTC: {Decimal: ".", Thousand: ",", Code: BTC, Fraction: 4, Grapheme: BTC, Template: "1 $", NetLevel: 1, NetID: "Bitcoin"},
	ETH: {Decimal: ".", Thousand: ",", Code: ETH, Fraction: 4, Grapheme: ETH, Template: "1 $", NetLevel: 1, NetID: "Etherium"},
}

func DefaultCrypto() Crypto {
	val, _ := FindByCode("BTC")

	return *val
}

func FindByCode(code string) (*Crypto, bool) {
	val, ok := cc[code]
	if !ok {
		empty := emptyCrypto()

		return &empty, false
	}

	return val, true
}

// AddCryptoCurrency lets you insert or update currency in currencies list.
func Add(c Crypto) bool {
	_, ok := cc[c.Code]
	if !ok {
		return ok
	}
	cc[c.Code] = &c

	return true
}

func (cr Crypto) Formatter() *f.Formatter {
	return &f.Formatter{
		Fraction: cr.Fraction,
		Decimal:  cr.Decimal,
		Thousand: cr.Thousand,
		Grapheme: cr.Grapheme,
		Template: cr.Template,
	}
}

func (cr Crypto) Format(amount decimal.Decimal) string {
	f := cr.Formatter()

	return f.Format(amount)
}

func (cr Crypto) IsFiat() bool {
	return false
}

type Crypto struct {
	Code     string
	Grapheme string
	Template string
	Decimal  string
	Thousand string
	NetID    string
	NetLevel int
	Fraction int
}

type Cryptos map[string]*Crypto

func emptyCrypto() Crypto {
	return Crypto{
		Code:     "",
		Grapheme: "",
		Template: "",
		Decimal:  "",
		Thousand: "",
		NetID:    "",
		NetLevel: -1,
		Fraction: 0,
	}
}

func New(crd CryptoRawData) Crypto {
	return Crypto{}
}
