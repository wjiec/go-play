package main

import (
	"go-play/xspew"
)

//go:generate go run go-play/wiring/a_quickstart

func main() {
	foobar, err := InitService("/etc/foo.conf", "/etc/bar.conf")
	if err != nil {
		panic(err)
	}

	xspew.Dump(foobar)
}
