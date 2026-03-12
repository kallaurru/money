package fiat

import (
	gm "github.com/Rhymond/go-money"
)

func New() *gm.Money {
	return gm.New(100, gm.RUR)
}
