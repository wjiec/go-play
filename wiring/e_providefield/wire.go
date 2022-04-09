//go:build wireinject

package main

import "github.com/google/wire"

func InitRedis(options *RedisConnectOption) (*Redis, error) {
	panic(wire.Build(
		NewRedis,
		// 当我们需要将一个结构体中的特定字段作为依赖项时，可以使用 wire.FieldsOf 方法
		// 这个方法需要传入指定的字段名字，不能使用 `*` 代替所有字段
		// 如果我们给定的结构体类型是一个指向指针的指针，那么 wire 可以提供字段的指针类型的依赖
		wire.FieldsOf(new(*RedisConnectOption), "Address", "Database", "Password"),
	))
}
