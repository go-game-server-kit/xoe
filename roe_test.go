package xoe

import (
	"fmt"
	"testing"
)

func TestRoe(t *testing.T) {
	// 不输出
	Roe(nil)
	// 输出
	Roe(fmt.Errorf("error1"))
	Roe(fmt.Errorf("error2"), WithStack(), WithData(Data{"key": "value"}), WithMsg("msg"))
}

func TestRoe1(t *testing.T) {
	e1 := func(msg string) (string, error) {
		if msg != "" {
			return "", fmt.Errorf(msg)
		}
		return "ok", nil
	}
	Roe1(e1(""))
	Roe1(e1("error1"))
	Roe1With(e1("error2"))(WithStack(), WithMsg("msg"), WithData(Data{"key": "value"}))
}

func TestRoe2(t *testing.T) {
	e2 := func(err string) (string, string, error) {
		if err != "" {
			return "", "", fmt.Errorf(err)
		}
		return "ok", "ok", nil
	}
	Roe2(e2(""))
	Roe2(e2("error1"))
	Roe2With(e2("error2"))(WithStack(), WithMsg("msg"), WithData(Data{"key": "value"}))
}
