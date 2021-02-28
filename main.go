package main

import (
	"home/config"
	"home/custom"
	"home/migration"
	"home/redis"
	"home/repository"
	"home/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	confErr, conf := config.Load()
	if confErr != nil {
		panic("can not init config")
	}

	repErr, rep := repository.InitRepository(conf)
	if repErr != nil {
		panic("can not init repository")
	}

	redisErr, rds := redis.InitRedis(conf)
	if redisErr != nil {
		panic("can not init redis")
	}

	cs := custom.NewCustom(rep, conf, rds)
	migration.MigrateDatabase(cs)

	router := gin.Default()
	routers.InitRouters(router, cs)

	router.Run(":7777")
}
