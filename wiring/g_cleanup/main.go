package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"go-play/xspew"
)

//go:generate go run go-play/wiring/g_cleanup

func main() {
	l, cleanup, err := InitService(Filepath("/tmp/rand-" + strconv.Itoa(rand.Int())))
	if err != nil {
		xspew.Dump(err)
		return
	}

	defer func() {
		cleanup()
		fmt.Println("cleanup service ...")
	}()

	xspew.Dump(l)
}
