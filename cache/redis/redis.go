package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

type RedisHelper struct {
	Self *redis.Client
}

var RH *RedisHelper

func (rh *RedisHelper) Init() {
	RH = &RedisHelper{
		Self: getSelfRedis(),
	}
}

func (rh *RedisHelper) Close() {
	_ = rh.Self.Close()
}

func getSelfRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})

}
