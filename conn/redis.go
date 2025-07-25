package conn

import (
	"github.com/hossainabid/go-ims/config"
	"github.com/hossainabid/go-ims/logger"

	"github.com/go-redis/redis"
)

var client *redis.Client

func ConnectRedis() {
	conf := config.Redis()

	logger.Info("connecting to redis at ", conf.Host, ":", conf.Port, "...")

	client = redis.NewClient(&redis.Options{
		Addr: conf.Host + ":" + conf.Port,
		DB:   conf.Db,
	})

	if _, err := client.Ping().Result(); err != nil {
		logger.Error("failed to connect redis: ", err)
		panic(err)
	}

	logger.Info("redis connection successful...")
}

func Redis() *redis.Client {
	return client
}
