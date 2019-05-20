package utils

import (
	"fmt"
	"milano.gaodun.com/pkg/consul"
	"time"

	"github.com/go-redis/redis"
)

type RedisClient struct {
	RedisClientHandle *redis.Client
}

var RedisHandle = initRedisClient()

func initRedisClient() *RedisClient {
	var redisStuct = new(RedisClient)
	redisStuct.RedisClientHandle = getClient()
	return redisStuct
}

func getClient() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     consul.GdConsul["PUBLIC_REDIS_HOST"] + ":" + consul.GdConsul["PUBLIC_REDIS_PORT"],
		Password: consul.GdConsul["PUBLIC_REDIS_PASSWD"], // no password set
		DB:       0,                                      // use default DB
	})
	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}

	return client
}

func (rc *RedisClient) SetData(key string, value interface{}, expiration time.Duration) {
	err := rc.RedisClientHandle.Set(key, value, expiration).Err()
	if err != nil {
		fmt.Println(err.Error())
	}
}
func (rc *RedisClient) Expire(key string, expiration time.Duration) {
	err := rc.RedisClientHandle.Expire(key, expiration).Err()
	if err != nil {
		fmt.Println(err.Error())
	}
}
func (rc *RedisClient) GetData(key string) interface{} {

	val, err := rc.RedisClientHandle.Get(key).Result()
	if err == redis.Nil {
		fmt.Println(key + " does not exist")
	} else if err != nil {
		fmt.Println(err.Error())
		// panic(err)
	}
	return val
}
