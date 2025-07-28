package services

import (
	"errors"

	"github.com/hossainabid/go-ims/consts"
	"github.com/hossainabid/go-ims/domain"
	"github.com/hossainabid/go-ims/types"
	"github.com/hossainabid/go-ims/utils/errutil"
)

type StockHistoryServiceImpl struct {
	stockHistoryRepo domain.StockHistoryRepository
	productRepo      domain.ProductRepository
}

func NewStockHistoryServiceImpl(stockHistoryRepo domain.StockHistoryRepository, productRepo domain.ProductRepository) *StockHistoryServiceImpl {
	return &StockHistoryServiceImpl{
		stockHistoryRepo: stockHistoryRepo,
		productRepo:      productRepo,
	}
}

func (svc *StockHistoryServiceImpl) RecordStockHistory(stockHistoryReq *types.RecordStockHistoryRequest) (*types.RecordStockHistoryResponse, error) {
	stockHistory := stockHistoryReq.ToStockHistory()
	product, err := svc.productRepo.ReadProductByID(stockHistory.ProductID)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errutil.ErrRecordNotFound
	}
	if stockHistory.OperationType == consts.OperationTypeRequisition {
		product.WarehouseQty += stockHistory.Qty
	} else if stockHistory.OperationType == consts.OperationTypePublishInLive {
		if stockHistory.Qty > product.WarehouseQty {
			return nil, errors.New("insufficient stock in warehouse to publish in live")
		} else {
			product.WarehouseQty -= stockHistory.Qty
			product.LiveQty += stockHistory.Qty
		}
	} else if stockHistory.OperationType == consts.OperationTypeRevertBackFromLive {
		if stockHistory.Qty > product.LiveQty {
			return nil, errors.New("insufficient stock in live to revert back from live")
		} else {
			product.LiveQty -= stockHistory.Qty
			product.WarehouseQty += stockHistory.Qty
		}
	} else if stockHistory.OperationType == consts.OperationTypeMarkDamage {
		if stockHistory.Qty > product.WarehouseQty {
			return nil, errors.New("insufficient stock in warehouse to mark damage")
		} else {
			product.WarehouseQty -= stockHistory.Qty
		}
	}
	createdStockHistory, err := svc.stockHistoryRepo.RecordStockHistory(stockHistory)
	if err != nil {
		return nil, err
	}
	_, err = svc.productRepo.UpdateProduct(product)
	if err != nil {
		return nil, err
	}

	return &types.RecordStockHistoryResponse{
		Message:      "Stock history recorded",
		StockHistory: createdStockHistory,
	}, nil
}

func (svc *StockHistoryServiceImpl) ListStockHistories(stockHistoryReq types.ListStockHistoryRequest) (*types.PaginatedStockHistoryResponse, error) {
	offset := (stockHistoryReq.Page - 1) * stockHistoryReq.Limit
	stockHistories, count, err := svc.stockHistoryRepo.ListStockHistories(stockHistoryReq.ProductID, stockHistoryReq.Limit, offset)
	if errors.Is(err, errutil.ErrRecordNotFound) {
		return &types.PaginatedStockHistoryResponse{}, nil
	}
	if err != nil {
		return nil, err
	}
	response := &types.PaginatedStockHistoryResponse{
		Page:           stockHistoryReq.Page,
		Limit:          stockHistoryReq.Limit,
		Total:          count,
		StockHistories: stockHistories,
	}
	return response, nil
}
