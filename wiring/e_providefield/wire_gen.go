// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func InitRedis(options *RedisConnectOption) (*Redis, error) {
	redisConnectString := options.Address
	redisDatabaseIndex := options.Database
	redisConnectPassword := options.Password
	redis, err := NewRedis(redisConnectString, redisDatabaseIndex, redisConnectPassword)
	if err != nil {
		return nil, err
	}
	return redis, nil
}