package worker

import (
	"fmt"
	"time"

	"github.com/hibiken/asynq"
	"github.com/hossainabid/go-ims/config"
	"github.com/hossainabid/go-ims/types"
)

func StartAsynqWorker(mux *asynq.ServeMux) {
	worker := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr: config.Asynq().RedisAddr,
			DB:   config.Asynq().DB,
		},
		asynq.Config{
			Concurrency: config.Asynq().Concurrency,
			Queues: map[string]int{
				config.Asynq().Queue: 1,
			},
			RetryDelayFunc: func(numOfRetry int, e error, t *asynq.Task) time.Duration {
				switch t.Type() {
				case types.AsynqTaskTypeStockSync.String():
					return config.Asynq().StockSyncTaskDelay * time.Second
				default:
					return asynq.DefaultRetryDelayFunc(numOfRetry, e, t)
				}
			},
		},
	)

	if err := worker.Run(mux); err != nil {
		panic(fmt.Sprintf("could not run worker: %v", err))
	}
}
