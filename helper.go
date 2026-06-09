package money

import (
	//	gm "github.com/Rhymond/go-money"
	"github.com/shopspring/decimal"
	"regexp"
)

func MakeAmountFromFloat64Exp(in float64, exp int32) decimal.Decimal {
	return decimal.NewFromFloatWithExponent(in, 4)
}

func MakeAmountFromValue(val int64, std bool) decimal.Decimal {
	if std {
		return decimal.New(val, 2)
	}

	return decimal.New(val, 4)
}

// MakeAmountFromSnakeFormat like a 1_000_000_000
func MakeAmountFromSnakeFormat(val string) (decimal.Decimal, error) {
	r := regexp.MustCompile("[_]")
	return decimal.NewFromFormattedString(val, r)
}

func MakePenny(val int64) decimal.Decimal {
	return NewPenny(val, true)
}

func MakePennyExtra(val int64) decimal.Decimal {
	return NewPenny(val, false)
}
