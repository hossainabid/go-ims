package domain

import (
	"github.com/hibiken/asynq"
	"github.com/hossainabid/go-ims/models"
	"github.com/hossainabid/go-ims/types"
)

type (
	AsynqRepository interface {
		CreateTask(event types.AsynqTaskType, payload interface{}) (*asynq.Task, error)
		EnqueueTask(task *asynq.Task, customOpts *types.AsynqOption) (string, error)
		DequeueTask(taskID string) error
	}

	AsynqService interface {
		CreateStockSyncTask(stockHistory *models.StockHistory) error
	}
)
