package money

import (
	gm "github.com/Rhymond/go-money"
)

type Crypto struct {
	_m       *gm.Money
	netID    string
	netLevel int
}

func New(m *gm.Money, netID string, netLevel int) *Crypto {
	return &Crypto{
		_m:       m,
		netID:    netID,
		netLevel: netLevel,
	}
}

func NewFromMoney(amount int64, code, netID string, netLevel int) *Crypto {
	return &Crypto{
		_m:       gm.New(amount, code),
		netID:    netID,
		netLevel: netLevel,
	}
}
