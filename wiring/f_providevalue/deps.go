package main

import (
	"fmt"
	"io"
)

type Message string

type Alerter interface {
	Alert()
}

type SimpleAlerter struct {
	Msg Message
	Out io.Writer
}

func (a *SimpleAlerter) Alert() {
	_, _ = fmt.Fprintln(a.Out, a.Msg)
}

func NewSimpleAlerter(out io.Writer, msg Message) *SimpleAlerter {
	return &SimpleAlerter{Out: out, Msg: msg}
}
