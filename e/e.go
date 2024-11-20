package e

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
)

type (
	WithAble interface {
		with(*XError)
	}
	XError struct {
		error
		data  map[string]any
		stack []uintptr
	}
	Data  map[string]any
	stack struct {
		IsOpen bool
	}
)

var (
	DefaultWithStack = true
)

func (e *XError) Unwrap() error {
	return e.error
}

func (d Data) with(e *XError) {
	if e.data == nil {
		e.data = make(Data)
	}
	for k, v := range d {
		e.data[k] = v
	}
}

func (s *stack) with(e *XError) {
	if s.IsOpen && e.stack == nil {
		e.stack = make([]uintptr, 64)
		n := runtime.Callers(3, e.stack)
		e.stack = e.stack[:n]
	}
}

func Stack(with bool) WithAble {
	return &stack{IsOpen: with}
}

func With(err error, args ...WithAble) error {
	if err == nil {
		return nil
	}
	// default with stack
	if DefaultWithStack {
		has := false
		for _, arg := range args {
			if _, ok := arg.(*stack); ok {
				has = true
				break
			}
		}
		if !has {
			args = append(args, Stack(true))
		}
	}
	if len(args) == 0 {
		return err
	}
	var e *XError
	if !errors.As(err, &e) {
		e = &XError{error: err}
	}
	for _, arg := range args {
		arg.with(e)
	}
	return e
}

func GetData(err error, key string) any {
	if err == nil {
		return nil
	}
	var e *XError
	if !errors.As(err, &e) {
		return nil
	}
	if e.data == nil {
		return nil
	}
	return e.data[key]
}

func GetAllData(err error) Data {
	if err == nil {
		return nil
	}
	var e *XError
	if !errors.As(err, &e) {
		return nil
	}
	return e.data
}

func GetStack(err error) []*runtime.Frame {
	if err == nil {
		return nil
	}
	var e *XError
	if !errors.As(err, &e) {
		return nil
	}

	if e.stack == nil {
		return nil
	}

	rs := make([]*runtime.Frame, 0, len(e.stack))
	frames := runtime.CallersFrames(e.stack)
	for {
		frame, more := frames.Next()
		rs = append(rs, &frame)
		if !more {
			break
		}
	}
	return rs
}

func FormatStack(stack []*runtime.Frame) []string {
	rs := make([]string, 0, len(stack))
	for _, frame := range stack {
		rs = append(rs, fmt.Sprintf("%s:%d %s", frame.File, frame.Line, frame.Function))
	}
	return rs
}

func ShortStack(stack []string) []string {
	pwd, _ := os.Getwd()
	pwd += string(os.PathSeparator)

	var rs []string
	for _, s := range stack {
		if strings.Contains(s, "/go/libexec/src/") {
			continue
		}

		if strings.Contains(s, pwd) {
			s = strings.Replace(s, pwd, "", 1)
		}

		rs = append(rs, s)
	}
	return rs
}
