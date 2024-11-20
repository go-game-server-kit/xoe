package xoe

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

type (
	Stack []uintptr
)

var (
	StackDepth = 32
)

func NewStack(skip int) Stack {
	pcs := make(Stack, StackDepth)
	n := runtime.Callers(1+skip, pcs[:])
	return pcs[:n]
}

func (s Stack) Frames() []*runtime.Frame {
	if len(s) == 0 {
		return nil
	}
	frames := runtime.CallersFrames(s)
	var fs []*runtime.Frame
	for {
		frame, more := frames.Next()
		fs = append(fs, &frame)
		if !more {
			break
		}
	}
	return fs
}

func (s Stack) ShortFile() string {
	if len(s) == 0 {
		return ""
	}
	frame, _ := runtime.CallersFrames(s).Next()
	return path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)
}

func (s Stack) Strings() []string {
	frames := s.Frames()
	if len(frames) == 0 {
		return nil
	}
	rs := make([]string, 0, len(frames))
	for _, frame := range frames {
		rs = append(rs, fmt.Sprintf("%s:%d %s", frame.File, frame.Line, frame.Function))
	}
	return rs
}

func (s Stack) String() string {
	return strings.Join(ShortStack(s.Strings()), "\n")
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
