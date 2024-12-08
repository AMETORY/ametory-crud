package database

import (
	"ametory-crud/config"
	"fmt"

	"github.com/go-redis/redis"
)

var REDIS = &redis.Client{}

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.App.Redis.Host, config.App.Redis.Port),
		Password: config.App.Redis.Password,
		DB:       config.App.Redis.DB,
	})
	REDIS = client
}
