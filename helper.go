package money

import (
	//	gm "github.com/Rhymond/go-money"
	cr "github.com/kallaurru/money/internal/crypto"
	crn "github.com/kallaurru/money/internal/currency"
	"github.com/shopspring/decimal"

	"regexp"
)

const (
	constFractionKop     = -2
	constFractionPens    = -4
	constFractionSatoshi = -8
)

func MakeCryptoAmount(code string) (decimal.Decimal, bool) {
	crypto, ok := cr.FindByCode(code)
	if !ok {
		return decimal.New(0, 0), ok
	}
	fraction := -1 * crypto.Fraction
	if fraction == 0 {
		return decimal.New(1, 0), true
	}

	if fraction >= constFractionKop {
		return MakeAmountAsKop(1), true
	}

	if fraction >= constFractionPens {
		return MakeAmountAsPens(1), true
	}

	return MakeAmountAsSatoshi(1), true
}

func MakeFiatAmount(code string) (decimal.Decimal, bool) {
	currency, ok := crn.FindByCode(code)
	if !ok {
		return decimal.New(0, 0), ok
	}

	fraction := currency.Fraction
	if fraction == 0 {
		return decimal.New(1, 0), true
	}

	if (-1 * fraction) >= constFractionKop {
		return MakeAmountAsKop(1), true
	}

	if (-1 * fraction) >= constFractionPens {
		return MakeAmountAsPens(1), true
	}

	return MakeAmountAsSatoshi(1), true
}

// Basic with exp -2
func MakeAmountAsKop(amount int64) decimal.Decimal {
	return decimal.New(amount, constFractionKop)
}

// Basic with exp -4
func MakeAmountAsPens(amount int64) decimal.Decimal {
	return decimal.New(amount, constFractionPens)
}

// Basic with exp -8
func MakeAmountAsSatoshi(amount int64) decimal.Decimal {
	return decimal.New(amount, constFractionSatoshi)
}

// MakeAmountFromSnakeFormat like a 1_000_000_000
func MakeAmountFromSnakeFormat(val string) (decimal.Decimal, error) {
	r := regexp.MustCompile("[_]")
	return decimal.NewFromFormattedString(val, r)
}
