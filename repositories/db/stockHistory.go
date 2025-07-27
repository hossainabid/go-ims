package db

import (
	"fmt"

	"github.com/hossainabid/go-ims/logger"
	"github.com/hossainabid/go-ims/models"
	"github.com/hossainabid/go-ims/utils/errutil"
)

func (repo *Repository) RecordStockHistory(stockHistory *models.StockHistory) (*models.StockHistory, error) {
	qry := repo.client.Create(stockHistory)
	if qry.Error != nil {
		logger.Error(fmt.Errorf("error creating stock history: %w", qry.Error))
		return nil, qry.Error
	}

	return stockHistory, nil
}

func (repo *Repository) ListStockHistories(product_id, limit, offset int) ([]*models.StockHistory, int, error) {
	var stockHistories []*models.StockHistory
	var count int64

	query := repo.client.Model(&models.StockHistory{}).Where("product_id =  ?", product_id)

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	result := query.Offset(offset).Limit(limit).Find(&stockHistories)
	if result.RowsAffected == 0 {
		logger.Error("no stock histories found")
		return nil, 0, errutil.ErrRecordNotFound
	}
	if result.Error != nil {
		logger.Error(fmt.Errorf("error listing stock histories: %w", result.Error))
		return nil, 0, result.Error
	}

	return stockHistories, int(count), nil
}
