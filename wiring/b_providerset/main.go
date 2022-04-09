package main

import "go-play/xspew"

//go:generate go run go-play/wiring/b_providerset

func main() {
	blog, err := InitService("mysql://root:root@tcp(127.0.0.1)/world?charset=utf8mb4")
	if err != nil {
		panic(err)
	}

	xspew.Dump(blog)
}
