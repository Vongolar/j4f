package jredis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var dbs = make(map[string]*redis.Client)

//GetDB 获取
func GetDB(name string) *redis.Client {
	return dbs[name]
}

//Connect 连接redis
func Connect(name string, addr string, password string, db int) error {
	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	cmd := c.Ping(context.Background())
	if err := cmd.Err(); err != nil {
		return err
	}
	dbs[name] = c
	return nil
}
