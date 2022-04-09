//go:build wireinject

package main

import (
	"io"

	"github.com/google/wire"
)

func InitLogger(out io.Writer) (Logger, error) {
	panic(wire.Build(
		NewConsoleLogger,
		// 由于我们需要一个 Logger 接口，但是我们提供的 ConsoleLogger
		// 返回的是一个具体的实例对象, wire 无法将其与 Logger 接口进行关联
		// 所以我们需要显式的将两个类型进行绑定
		wire.Bind(new(Logger), new(*Console)),
	))
}
