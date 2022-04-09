package main

import "github.com/go-redis/redis/v8"

type RedisConnectString string

type RedisDatabaseIndex int

type RedisConnectPassword string

type RedisConnectOption struct {
	Address  RedisConnectString
	Database RedisDatabaseIndex
	Password RedisConnectPassword
}

type Redis struct {
	rdb *redis.Client
}

func NewRedis(addr RedisConnectString, db RedisDatabaseIndex, password RedisConnectPassword) (*Redis, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     string(addr),
		DB:       int(db),
		Password: string(password),
	})

	return &Redis{rdb: rdb}, nil
}
