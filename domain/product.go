package domain

import (
	"github.com/hossainabid/go-ims/models"
	"github.com/hossainabid/go-ims/types"
)

type (
	ProductService interface {
		CreateProduct(productReq *types.CreateProductRequest) (*types.CreateProductResponse, error)
		ListProducts(productReq types.ListProductRequest) (*types.PaginatedProductResponse, error)
		ReadProductByID(id int) (*models.Product, error)
		UpdateProduct(productReq *types.UpdateProductRequest) (*types.UpdateProductResponse, error)
		DeleteProduct(id int) (*types.DeleteProductResponse, error)
	}

	ProductRepository interface {
		CreateProduct(product *models.Product) (*models.Product, error)
		ListProducts(limit, offset int) ([]*models.Product, int, error)
		ReadProductByID(id int) (*models.Product, error)
		UpdateProduct(product *models.Product) (*models.Product, error)
		DeleteProduct(id int) error
	}
)
