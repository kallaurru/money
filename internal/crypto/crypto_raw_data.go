package crypto

import (
	"fmt"
	"strconv"
	"strings"
)

type CryptoRawData struct {
	_delim   string
	Fraction int
	NetLevel int
	Code     string
	NetID    string
	Grapheme string
	Template string
	Decimal  string
	Thousand string
}

func FromCSV(data, delim string) (CryptoRawData, error) {
	limit := 2 // required fraction and code properties
	parts := strings.Split(data, delim)
	if len(parts) < limit {
		return CryptoRawData{}, errMsgBadLine(limit)
	}

	fract := parts[0]
	code := parts[1]

	fractInt, err := strconv.Atoi(fract)
	if err != nil {
		return CryptoRawData{}, err
	}
	if code == "" {
		return CryptoRawData{}, errMsgNullObject()
	}
	obj := CryptoRawData{
		Fraction: fractInt,
		Code:     code,
	}
	if len(parts) > 2 {
		val, err := strconv.Atoi(parts[2])
		if err != nil {
			return CryptoRawData{}, err
		}
		obj.NetLevel = val
	}

	if len(parts) > 3 {
		obj.NetID = parts[3]
	}
	if len(parts) > 4 {
		obj.NetID = parts[4]
	}
	if len(parts) > 5 {
		obj.Decimal = parts[5]
	}
	if len(parts) > 6 {
		obj.Thousand = parts[6]
	}
	if len(parts) > 7 {
		obj.Grapheme = parts[7]
	}

	return obj, nil
}

func FromCrypto(c Crypto, delim string) CryptoRawData {
	return CryptoRawData{
		_delim:   delim,
		Fraction: c.Fraction,
		NetLevel: c.NetLevel,
		Code:     c.Code,
		NetID:    c.NetID,
		Grapheme: c.Grapheme,
		Template: c.Template,
		Decimal:  c.Decimal,
		Thousand: c.Thousand,
	}
}

// String - сериализация в строку. Здесь указана последовательность полей для сохранения
func (crd CryptoRawData) String() string {
	return fmt.Sprintf("%d%s%s%s%d%s%s%s%s%s%s%s%s%s%s\n",
		crd.Fraction, crd._delim,
		crd.Code, crd._delim,
		crd.NetLevel, crd._delim,
		crd.NetID, crd._delim,
		crd.Decimal, crd._delim,
		crd.Thousand, crd._delim,
		crd.Template, crd._delim,
		crd.Grapheme)
}
