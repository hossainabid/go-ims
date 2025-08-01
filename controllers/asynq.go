package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/hossainabid/go-ims/consts"
	"github.com/hossainabid/go-ims/domain"
	"github.com/hossainabid/go-ims/logger"
	"github.com/hossainabid/go-ims/models"
	"github.com/hossainabid/go-ims/utils/errutil"
)

type AsynqController struct {
	productRepo domain.ProductRepository
}

func NewAsynqController(productRepo domain.ProductRepository) *AsynqController {
	return &AsynqController{
		productRepo: productRepo,
	}
}

func (ac *AsynqController) ProcessStockSyncTask(ctx context.Context, t *asynq.Task) (err error) {
	logger.Info(fmt.Sprintf("Received task event [%s] with ID [%s]", t.Type(), t.ResultWriter().TaskID()))
	var stockHistory models.StockHistory

	if err = json.Unmarshal(t.Payload(), &stockHistory); err != nil {
		logger.Error(err)
		return
	}
	product, err := ac.productRepo.ReadProductByID(stockHistory.ProductID)
	if err != nil {
		return err
	}
	if product == nil {
		return errutil.ErrRecordNotFound
	}
	if stockHistory.OperationType == consts.OperationTypeRequisition {
		product.WarehouseQty += stockHistory.Qty
	} else if stockHistory.OperationType == consts.OperationTypePublishInLive {
		if stockHistory.Qty > product.WarehouseQty {
			return errors.New("insufficient stock in warehouse to publish in live")
		} else {
			product.WarehouseQty -= stockHistory.Qty
			product.LiveQty += stockHistory.Qty
		}
	} else if stockHistory.OperationType == consts.OperationTypeRevertBackFromLive {
		if stockHistory.Qty > product.LiveQty {
			return errors.New("insufficient stock in live to revert back from live")
		} else {
			product.LiveQty -= stockHistory.Qty
			product.WarehouseQty += stockHistory.Qty
		}
	} else if stockHistory.OperationType == consts.OperationTypeMarkDamage {
		if stockHistory.Qty > product.WarehouseQty {
			return errors.New("insufficient stock in warehouse to mark damage")
		} else {
			product.WarehouseQty -= stockHistory.Qty
		}
	}

	_, err = ac.productRepo.UpdateProduct(product)
	if err != nil {
		return err
	}
	t.ResultWriter().Write([]byte(fmt.Sprintf("Stock sync task created successfully for product id: %d", stockHistory.ProductID)))

	return
}
