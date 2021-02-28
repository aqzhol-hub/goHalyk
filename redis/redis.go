package redis

import (
	"github.com/go-redis/redis"
)

type RedisClient interface {
	SaveAuth(signedToken string, userID uint) error
	GetAuth(signedToken string) (uint, error)
	RemoveAuth(signedToken string) error
}

type redisClient struct {
	client *redis.Client
}
