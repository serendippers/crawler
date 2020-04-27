package initialize

import (
	"crawler/global"
	"github.com/go-redis/redis"
)

func Redis() {

	redisConfig := global.CONFIG.Redis

	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})

	pong, err := client.Ping().Result()

	if err != nil {
		global.LOG.Error(err)
	} else {
		global.LOG.Info("redis connect ping response:", pong)
		global.REDIS = client
	}
}
