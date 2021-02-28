package redis_test

import (
	"home/config"
	"home/redis"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitRedis(t *testing.T) {

	conf := &config.Config{
		Redis: struct {
			Network  string `default:"tcp"`
			Addr     string `default:"localhost:6379"`
			Password string `default:""`
			DB       int    `default:"0"`
		}{
			Network:  "tcp",
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	}

	err, c := redis.InitRedis(conf)
	assert.NoError(t, err)

	err = c.SaveAuth("qweqwe", 1)
	assert.NoError(t, err)

	id, err2 := c.GetAuth("qweqwe")
	assert.NoError(t, err2)
	assert.Equal(t, uint(1), id)

	err = c.RemoveAuth("qweqwe")
	assert.NoError(t, err)

}
