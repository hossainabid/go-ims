package controllers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/hossainabid/go-ims/domain"
	"github.com/hossainabid/go-ims/logger"
	"github.com/hossainabid/go-ims/models"
)

type AsynqController struct {
	productSvc domain.ProductService
	mailSvc    domain.MailService
}

func NewAsynqController(productSvc domain.ProductService, mailSvc domain.MailService) *AsynqController {
	return &AsynqController{
		productSvc: productSvc,
		mailSvc:    mailSvc,
	}
}

func (ac *AsynqController) ProcessStockSyncTask(ctx context.Context, t *asynq.Task) (err error) {
	logger.Info(fmt.Sprintf("Received task event [%s] with ID [%s]", t.Type(), t.ResultWriter().TaskID()))
	var stockHistory models.StockHistory

	if err = json.Unmarshal(t.Payload(), &stockHistory); err != nil {
		logger.Error(err)
		return
	}

	err = ac.productSvc.StockSync(stockHistory)
	if err != nil {
		return err
	}

	err = ac.mailSvc.SendLowStockEmail(stockHistory.ProductID)
	if err != nil {
		return err
	}

	t.ResultWriter().Write([]byte(fmt.Sprintf("Stock sync task created successfully for product id: %d", stockHistory.ProductID)))

	return nil
}
