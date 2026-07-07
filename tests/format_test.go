package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/kallaurru/money"
)

func TestFormatStringCrypto(t *testing.T) {
	ethInPens := money.MakeAmountAsPens(11636310)
	crypto, ok := money.NewCrypto(ethInPens, "ETH")
	if !ok {
		t.Errorf("bad currency code")
		os.Exit(1)
	}
	fmt.Printf("Simple Amount - %s\n", ethInPens.String())
	fmt.Printf("Amount - %s", crypto.Format())
}

func TestFormatStringFiat(t *testing.T) {
	rubInKop := money.MakeAmountAsKop(10000)
	rub, ok := money.New(rubInKop, "RUB")

	if !ok {
		t.Errorf("bad currency code RUB")
		os.Exit(1)
	}

	fmt.Printf("Simple Amount - %s\n", rub.String())
	fmt.Printf("Amount - %s", rub.Format())
}
