package main

import (
	"fmt"
	"io"
)

type Logger interface {
	Log(msg string, args ...interface{})
}

type Console struct {
	out io.Writer
}

func (log *Console) Log(msg string, args ...interface{}) {
	_, _ = fmt.Fprintf(log.out, ">> %s [%+v]", msg, args)
}

func NewConsoleLogger(out io.Writer) *Console {
	return &Console{out: out}
}
