package xoe

import (
	"fmt"
	"testing"
)

func TestLoe(t *testing.T) {
	// 不输出
	Loe(nil)
	// 输出
	Loe(fmt.Errorf("error1"))
	Loe(fmt.Errorf("error2"), WithStack(), WithData(Data{"key": "value"}), WithMsg("msg"))
}

func TestLoe1(t *testing.T) {
	e1 := func(msg string) (string, error) {
		if msg != "" {
			return "", fmt.Errorf(msg)
		}
		return "ok", nil
	}
	Loe1(e1(""))
	Loe1(e1("error1"))
	Loe1With(e1("error2"))(WithStack(), WithMsg("msg"), WithData(Data{"key": "value"}))
}

func TestLoe2(t *testing.T) {
	e2 := func(err string) (string, string, error) {
		if err != "" {
			return "", "", fmt.Errorf(err)
		}
		return "ok", "ok", nil
	}
	Loe2(e2(""))
	Loe2(e2("error1"))
	Loe2With(e2("error2"))(WithStack(), WithMsg("msg"), WithData(Data{"key": "value"}))
}
