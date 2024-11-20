package xoe

import (
	"log"
	"os"
)

type (
	Logger func(err error, arg *Arg)
)

var (
	_logger = log.New(os.Stderr, "[xoe:log] ", log.LstdFlags)
	logger  = func(err error, arg *Arg) {
		_logger.Println(NewStack(arg.StackSkip+4).ShortFile(), arg.String(err))
	}
)

func SetLogger(l Logger) {
	logger = l
}

func Log(tag string, err error, args []Args) {
	arg := NewArg(tag, args)
	arg.Logger(err, arg)
}
