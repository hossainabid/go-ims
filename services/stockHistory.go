package services

import (
	"errors"

	"github.com/hossainabid/go-ims/domain"
	"github.com/hossainabid/go-ims/types"
	"github.com/hossainabid/go-ims/utils/errutil"
)

type StockHistoryServiceImpl struct {
	stockHistoryRepo domain.StockHistoryRepository
}

func NewStockHistoryServiceImpl(stockHistoryRepo domain.StockHistoryRepository) *StockHistoryServiceImpl {
	return &StockHistoryServiceImpl{
		stockHistoryRepo: stockHistoryRepo,
	}
}

func (svc *StockHistoryServiceImpl) RecordStockHistory(stockHistoryReq *types.RecordStockHistoryRequest) (*types.RecordStockHistoryResponse, error) {
	stockHistory := stockHistoryReq.ToStockHistory()
	createdStockHistory, err := svc.stockHistoryRepo.RecordStockHistory(stockHistory)
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
