package xoe

import (
	"fmt"
	"github.com/go-game-server-kit/xoe/e"
	"strings"
	"testing"
)

func TestLoe(t *testing.T) {
	Loe(nil)
	Loe(fmt.Errorf("error"))
	Loe1(1, nil)
	Loe1(1, fmt.Errorf("error"))
}

func TestRoe(t *testing.T) {
	Roe(nil)
	Roe(fmt.Errorf("error"))
	Roe1(1, nil)
	Roe1(1, fmt.Errorf("error"))
}

func TestPoe(t *testing.T) {
	Poe(nil)
	Poe(fmt.Errorf("error"))
}

func TestPoe1(t *testing.T) {
	Poe1(1, nil)
	Poe1(1, fmt.Errorf("error"))
}

func TestData(t *testing.T) {
	err := fmt.Errorf("error")
	err = e.With(err, e.Data{"key": "value"})
	fmt.Println(e.GetData(err, "key"))
	fmt.Println(e.GetAllData(err))
}

func TestStacks(t *testing.T) {
	err := e.With(fmt.Errorf("error"), e.Stack(true))
	fmt.Println(strings.Join(e.FormatStack(e.GetStack(err)), "\n"))
}

func TestStacks1(t *testing.T) {
	e.DefaultWithStack = true
	err := e.With(fmt.Errorf("error"))
	fmt.Println(strings.Join(e.FormatStack(e.GetStack(err)), "\n"))
}

func TestStacks2(t *testing.T) {
	e.DefaultWithStack = true
	err := e.With(fmt.Errorf("error"), e.Stack(false))
	fmt.Println(strings.Join(e.FormatStack(e.GetStack(err)), "\n"))
}
