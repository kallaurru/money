package crypto

import (
	"fmt"
)

const (
	ErrMsgBadLine = ""
)

func errMsgBadLine(params int) error {
	return fmt.Errorf("bad line. Count of params shoud be equal - %d", params)
}

func errMsgNullObject() error {
	return fmt.Errorf("object is empty. Code is empty")
}
