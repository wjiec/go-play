//go:build wireinject

package main

import "github.com/google/wire"

func InitService(dsn MySQLConnectString) (*BlogService, error) {
	panic(wire.Build(
		NewBlogService,
		NewUserRepository,
		NewArticleRepository,
		// 引入一个集合
		MySQLProviderSet,
	))
}
