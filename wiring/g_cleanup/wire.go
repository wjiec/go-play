//go:build wireinject

package main

import (
	"github.com/google/wire"
)

func InitService(file Filepath) (*HalfLiveCat, func(), error) {
	panic(wire.Build(
		NewFileLogger,
		NewHalfLiveCat,
	))
}
