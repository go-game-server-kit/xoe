package xoe

import (
	"errors"
	"fmt"
	"testing"
)

func TestEoe(t *testing.T) {
	Loe(a())
	Roe(a())
	Poe(a())
}

var _err = &_strErr{_strErr{fmt.Errorf("error")}}

type _strErr struct {
	error
}

func TestEoeIs(t *testing.T) {
	err := Eoe(_err)
	if !errors.Is(err, _err) {
		t.Error("EoeIs failed")
	}
	var err1 *_strErr
	if !errors.As(err, &err1) {
		t.Error("EoeAs failed")
	}
	if !errors.Is(errors.Unwrap(err), _err) {
		t.Error("EoeUnwrap failed")
	}
}

func a() error {
	return b()
}

func b() error {
	return Eoe(fmt.Errorf("error1"))
}
