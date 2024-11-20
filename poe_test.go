package xoe

import (
	"fmt"
	"testing"
)

func TestPoe(t *testing.T) {
	// 不输出
	Poe(nil)
	// 输出
	Poe(fmt.Errorf("error1"))
}

func TestPoeWith(t *testing.T) {
	Poe(fmt.Errorf("error2"), WithStack(), WithData(Data{"key": "value"}), WithMsg("msg"))
}

func TestPoe1(t *testing.T) {
	e1 := func(msg string) (string, error) {
		if msg != "" {
			return "", fmt.Errorf(msg)
		}
		return "ok", nil
	}
	Poe1(e1(""))
	Poe1(e1("error1"))
}

func TestPoe2(t *testing.T) {
	e2 := func(err string) (string, string, error) {
		if err != "" {
			return "", "", fmt.Errorf(err)
		}
		return "ok", "ok", nil
	}
	Poe2(e2(""))
	Poe2With(e2("error2"))(WithStack(), WithMsg("msg"), WithData(Data{"key": "value"}))
}
