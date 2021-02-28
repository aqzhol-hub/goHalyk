package custom

import (
	"home/config"
	"home/repository"

	"home/redis"
)

type Custom interface {
	GetRepository() repository.Repository
	GetConfig() *config.Config
	GetRedisClient() redis.RedisClient
}

type custom struct {
	rep    repository.Repository
	conf   *config.Config
	client redis.RedisClient
}

func NewCustom(rep repository.Repository, conf *config.Config, client redis.RedisClient) Custom {
	return &custom{rep: rep, conf: conf, client: client}
}

func (c *custom) GetRepository() repository.Repository {
	return c.rep
}

func (c *custom) GetConfig() *config.Config {
	return c.conf
}

func (c *custom) GetRedisClient() redis.RedisClient {
	return c.client
}
