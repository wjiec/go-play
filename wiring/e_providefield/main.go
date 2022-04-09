package main

import "go-play/xspew"

//go:generate go run go-play/wiring/e_providefield

func main() {
	rdb, err := InitRedis(&RedisConnectOption{})
	if err != nil {
		panic(err)
	}

	xspew.Dump(rdb)
}
