package domain

import (
	"github.com/hossainabid/go-ims/models"
	"github.com/hossainabid/go-ims/types"
)

type (
	StockHistoryService interface {
		RecordStockHistory(stockHistoryReq *types.RecordStockHistoryRequest) (*types.RecordStockHistoryResponse, error)
		ListStockHistories(stockHistoryReq types.ListStockHistoryRequest) (*types.PaginatedStockHistoryResponse, error)
	}

	StockHistoryRepository interface {
		RecordStockHistory(stockHistory *models.StockHistory) (*models.StockHistory, error)
		ListStockHistories(product_id, limit, offset int) ([]*models.StockHistory, int, error)
	}
)
