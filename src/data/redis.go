package data

import (
	"fmt"
	"github.com/aminsaki/golang-clean-web-api/config"
	"github.com/go-redis/redis/v7"
	"time"
)

var rediClient *redis.Client

func InitRedis(cfg *config.Config) {

	rediClient = redis.NewClient(&redis.Options{
		Addr:               fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		password:           cfg.Redis.Password,
		db:                 0,
		dialTimeout:        cfg.Redis.DialTimeout * time.Second,
		readTimeout:        cfg.Redis.ReadTimeout * time.Second,
		writeTimeout:       cfg.Redis.WriteTimeout * time.Second,
		poolSize:           cfg.Redis.PoolSize,
		poolTimeout:        cfg.Redis.PoolTimeout,
		IdleCheckFrequency: cfg.Redis.PoolTimeout,
	})

}

func GetRedis() *redis.Client {
	return rediClient
}

func CloseRedis() {

	rediClient.Close()

}
