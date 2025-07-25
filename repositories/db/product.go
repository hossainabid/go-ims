package db

import (
	"errors"
	"fmt"

	"github.com/hossainabid/go-ims/logger"
	"github.com/hossainabid/go-ims/models"
	"github.com/hossainabid/go-ims/types"
	"github.com/hossainabid/go-ims/utils/errutil"
	"gorm.io/gorm"
)

func (repo *Repository) CreateProduct(product *models.Product) (*models.Product, error) {
	qry := repo.client.Create(product)
	if qry.Error != nil {
		logger.Error(fmt.Errorf("error creating product: %w", qry.Error))
		return nil, qry.Error
	}

	return product, nil
}

func (repo *Repository) ListProducts(filter *types.ProductFilter, Limit, Offset int) ([]*models.Product, int, error) {
	var products []*models.Product
	var count int64

	query := repo.client.Model(&models.Product{})
	repo.applyFilters(query, filter)

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	result := query.Offset(Offset).Limit(Limit).Find(&products)
	if result.RowsAffected == 0 {
		logger.Error("no products found")
		return nil, 0, errutil.ErrRecordNotFound
	}
	if result.Error != nil {
		logger.Error(fmt.Errorf("error listing products: %w", result.Error))
		return nil, 0, result.Error
	}

	return products, int(count), nil
}

func (repo *Repository) applyFilters(query *gorm.DB, filter *types.ProductFilter) {
	if filter == nil {
		return
	}
	if filter.CreatedBy != nil {
		query = query.Where("created_by = ?", filter.CreatedBy)
	}
}

func (repo *Repository) ReadProductByID(id int) (*models.Product, error) {
	var product models.Product
	qry := repo.client.First(&product, id)
	if errors.Is(qry.Error, gorm.ErrRecordNotFound) {
		logger.Error(fmt.Errorf("product with ID %d not found", id))
		return nil, errutil.ErrRecordNotFound
	}
	if qry.Error != nil {
		logger.Error(fmt.Errorf("error getting product by ID: %w", qry.Error))
		return nil, qry.Error
	}

	return &product, nil
}

func (repo *Repository) UpdateProduct(product *models.Product) (*models.Product, error) {
	qry := repo.client.Where("id = ?", product.ID).Updates(product)
	if errors.Is(qry.Error, gorm.ErrRecordNotFound) {
		logger.Error(fmt.Errorf("no product found with ID %d", product.ID))
		return nil, errutil.ErrRecordNotFound
	}
	if qry.Error != nil {
		logger.Error(fmt.Errorf("error updating product: %w", qry.Error))
		return nil, qry.Error
	}
	return product, nil
}

func (repo *Repository) DeleteProduct(id int) error {
	qry := repo.client.Where("id = ?", id).Delete(&models.Product{})
	if qry.RowsAffected == 0 {
		logger.Error(fmt.Errorf("no product found with ID %d", id))
		return errutil.ErrRecordNotFound
	}
	if qry.Error != nil {
		logger.Error(fmt.Errorf("error deleting product: %w", qry.Error))
		return qry.Error
	}
	return nil
}
