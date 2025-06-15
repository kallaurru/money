package money

import (
	gm "github.com/Rhymond/go-money"
	"strings"
)

type CryptoCurrency struct {
	Code     string
	Template string
	Decimal  string
	Thousand string
	NetID    string
	NetLevel int
	Fraction int
}

type CC map[string]*CryptoCurrency

var cc = CC{
	BTC: {Decimal: ".", Thousand: ",", Code: "", Fraction: 4, Template: "1 BTC"},
	ETH: {Decimal: ".", Thousand: ",", Code: "", Fraction: 2, Template: "1 ETH"},
}

func (cc CC) Add(c *CryptoCurrency) CC {
	cc[c.Code] = c
	return cc
}

// CurrencyByCode returns the currency given the currency code defined as a constant.
func (cc CC) CurrencyByCode(code string) *CryptoCurrency {
	sc, ok := cc[code]
	if !ok {
		return nil
	}

	return sc
}

// AddCryptoCurrency lets you insert or update currency in currencies list.
func AddCryptoCurrency(code, Grapheme, Template, Decimal, Thousand string, Fraction int) *CryptoCurrency {
	c := CryptoCurrency{
		Code:     code,
		Template: Template,
		Decimal:  Decimal,
		Thousand: Thousand,
		Fraction: Fraction,
	}
	cc.Add(&c)
	return &c
}

func newCryptoCurrency(code string) *CryptoCurrency {
	return &CryptoCurrency{Code: strings.ToUpper(code)}
}

// GetCryptoCurrency returns the currency given the code.
func GetCryptoCurrency(code string) *CryptoCurrency {
	return cc.CurrencyByCode(strings.ToUpper(code))
}

// Formatter returns currency formatter representing
// used currency structure.
func (cr *CryptoCurrency) Formatter() *gm.Formatter {
	return &gm.Formatter{
		Fraction: cr.Fraction,
		Decimal:  cr.Decimal,
		Thousand: cr.Thousand,
		Grapheme: "",
		Template: cr.Template,
	}
}

// getDefault represent default currency if currency is not found in currencies list.
// Grapheme and Code fields will be changed by currency code.
func (cr *CryptoCurrency) getDefault() *CryptoCurrency {
	return &CryptoCurrency{Decimal: ".", Thousand: ",", Code: cr.Code, Fraction: 2, Template: "1 RUB"}
}

// get extended currency using currencies list.
func (cr *CryptoCurrency) get() *CryptoCurrency {
	if curr, ok := cc[cr.Code]; ok {
		return curr
	}

	return cr.getDefault()
}

func (cr *CryptoCurrency) equals(oc *CryptoCurrency) bool {
	return cr.Code == oc.Code
}
