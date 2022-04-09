package main

import (
	"bytes"

	"go-play/xspew"
)

//go:generate go run go-play/wiring/c_bindinterface

func main() {
	var buf bytes.Buffer
	logger, err := InitLogger(&buf)
	if err != nil {
		panic(err)
	}

	xspew.Dump(logger)
}
