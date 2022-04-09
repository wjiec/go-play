package main

import "go-play/xspew"

//go:generate go run go-play/wiring/d_injectstruct

func main() {
	xspew.Dump(InitCombine())
}
