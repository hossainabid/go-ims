package conn

import (
	"github.com/hibiken/asynq"
	"github.com/hossainabid/go-ims/config"
)

var asyncClient *asynq.Client
var asynqInspector *asynq.Inspector

func InitAsynqClient() {
	asyncClient = asynq.NewClient(asynq.RedisClientOpt{
		Addr: config.Asynq().RedisAddr,
		DB:   config.Asynq().DB,
	})
}

func InitAsyncInspector() {
	asynqInspector = asynq.NewInspector(asynq.RedisClientOpt{
		Addr: config.Asynq().RedisAddr,
		DB:   config.Asynq().DB,
	})
}

func Asynq() *asynq.Client {
	return asyncClient
}

func AsynqInspector() *asynq.Inspector {
	return asynqInspector
}
