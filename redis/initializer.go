package redis

import (
	"home/config"

	"github.com/go-redis/redis"
)

func InitRedis(conf *config.Config) (error, RedisClient) {
	client := redis.NewClient(&redis.Options{
		Network:  conf.Redis.Network,
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return err, nil
	}

	return nil, &redisClient{client: client}
}
