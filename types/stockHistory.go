package types

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hossainabid/go-ims/models"
)

type (
	RecordStockHistoryRequest struct {
		ProductID     int    `json:"product_id"`
		Qty           int    `json:"qty"`
		OperationType string `json:"operation_type"`
		Operation     string `json:"operation"`
		CreatedBy     int    `json:"created_by"`
	}

	RecordStockHistoryResponse struct {
		Message      string               `json:"message"`
		StockHistory *models.StockHistory `json:"stock_history"`
	}

	ListStockHistoryRequest struct {
		ProductID int `query:"product_id"`
		Page      int `query:"page"`
		Limit     int `query:"limit"`
	}

	PaginatedStockHistoryResponse struct {
		Total          int                    `json:"total"`
		Page           int                    `json:"page"`
		Limit          int                    `json:"limit"`
		StockHistories []*models.StockHistory `json:"stock_histories"`
	}
)

func (rshreq *RecordStockHistoryRequest) Validate() error {
	return v.ValidateStruct(rshreq,
		v.Field(&rshreq.ProductID, v.Required),
		v.Field(&rshreq.Qty, v.Required),
		v.Field(&rshreq.OperationType, v.Required),
		v.Field(&rshreq.Operation, v.Required),
	)
}

func (rshreq *RecordStockHistoryRequest) ToStockHistory() *models.StockHistory {
	stockHistory := &models.StockHistory{
		ProductID:     rshreq.ProductID,
		Qty:           rshreq.Qty,
		OperationType: rshreq.OperationType,
		Operation:     rshreq.Operation,
		CreatedBy:     rshreq.CreatedBy,
	}
	return stockHistory
}

func (lshreq *ListStockHistoryRequest) Validate() error {
	return v.ValidateStruct(lshreq,
		v.Field(&lshreq.ProductID, v.Required),
	)
}
