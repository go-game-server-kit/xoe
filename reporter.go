package xoe

import (
	"log"
	"os"
)

type (
	Reporter func(err error, arg *Arg)
)

var (
	_reporter = log.New(os.Stderr, "[xoe:report] ", log.LstdFlags)
	reporter  = func(err error, arg *Arg) {
		_reporter.Println(NewStack(arg.StackSkip+4).ShortFile(), arg.String(err))
	}
)

func SetReporter(r Reporter) {
	reporter = r
}

func Report(tag string, err error, args []Args) {
	arg := NewArg(tag, args)
	arg.Reporter(err, arg)
}
