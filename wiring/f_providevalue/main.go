package main

import "go-play/xspew"

//go:generate go run go-play/wiring/f_providevalue

func main() {
	xspew.Dump(NewAlerter())
}
