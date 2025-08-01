package services

import (
	"errors"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/hossainabid/go-ims/config"
	"github.com/hossainabid/go-ims/domain"
	"github.com/hossainabid/go-ims/logger"
	"github.com/hossainabid/go-ims/models"
	"github.com/hossainabid/go-ims/types"
)

type AsynqService struct {
	config    *config.AsynqConfig
	asynqRepo domain.AsynqRepository
}

func NewAsynqService(
	config *config.AsynqConfig,
	asynqRepo domain.AsynqRepository,
) *AsynqService {
	return &AsynqService{
		config:    config,
		asynqRepo: asynqRepo,
	}
}

func (svc *AsynqService) CreateStockSyncTask(stockHistory *models.StockHistory) error {

	taskID := fmt.Sprintf("%s_stock_history:%d", types.AsynqTaskTypeStockSync, stockHistory.ID)
	customOpts := &types.AsynqOption{
		Queue:        svc.config.Queue,
		TaskID:       taskID,
		DelaySeconds: svc.config.StockSyncTaskDelay,
		Retry:        svc.config.RetryCount,
	}

	task, err := svc.asynqRepo.CreateTask(types.AsynqTaskTypeStockSync, stockHistory)
	if err != nil {
		logger.Error(fmt.Sprintf("error: [%v] occurred while syncing stock task for stock history id: %d", err, stockHistory.ID))
		return err
	}

	_, err = svc.enqueueTask(task, customOpts)
	if err != nil {
		logger.Error(fmt.Sprintf("error: [%v] occurred while syncing stock task for stock history id: %d", err, stockHistory.ID))
		return err
	}
	return nil
}

func (svc *AsynqService) enqueueTask(task *asynq.Task, customOpts *types.AsynqOption) (taskID string, err error) {
	err = svc.asynqRepo.DequeueTask(customOpts.TaskID) // Ensure no duplicate tasks
	if err != nil && !errors.Is(err, asynq.ErrTaskNotFound) {
		logger.Error(fmt.Sprintf("error: [%v] occurred while dequeuing task with ID: %s", err, customOpts.TaskID))
	}

	taskID, err = svc.asynqRepo.EnqueueTask(task, customOpts)
	if errors.Is(err, asynq.ErrDuplicateTask) {
		logger.Warn(fmt.Sprintf("skipped: duplicate task for taskID: [%s]", customOpts.TaskID))
		err = nil // No error for duplicate tasks, just skip
		return
	}
	if err != nil {
		logger.Error(fmt.Sprintf("error: [%v] occurred while enqueuing task with ID: %s", err, customOpts.TaskID))
		return
	}

	logger.Info(fmt.Sprintf("enqueued task [%s] successfully", taskID))
	return taskID, nil
}
