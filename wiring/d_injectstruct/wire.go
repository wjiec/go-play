//go:build wireinject

package main

import "github.com/google/wire"

func InitCombine() *Service {
	panic(wire.Build(
		NewFoo,
		NewBar,
		NewFoobar,
		// 当我们需要多个依赖项包装在一个 struct 中时，可以直接通过 wire.Struct
		// 并通过传入需要注入的字段名称列表(*为全部注入)来达到这个效果，省去了再写一个
		// NewXXX(a A, b B, c C) *Combine {} 的方法
		wire.Struct(new(Service), "*"),
	))
}
