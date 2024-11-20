package xoe

import (
	"log"
	"os"
)

type (
	Reporter func(error)
)

var (
	logger   = log.New(os.Stderr, "[xoe] ", log.LstdFlags|log.Lshortfile)
	reporter = func(err error) {
		logger.Println("Report:", err)
	}
)

func SetLogger(l *log.Logger) {
	logger = l
}

func SetReporter(r Reporter) {
	reporter = r
}
