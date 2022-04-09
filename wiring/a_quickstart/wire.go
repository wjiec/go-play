//go:build wireinject

package main

import "github.com/google/wire"

func InitService(cfg1 FooConfig, cfg2 BarConfig) (*FooBar, error) {
	panic(wire.Build(NewFoo, NewBar, NewFooBar))
}
