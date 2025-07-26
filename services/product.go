package services

import (
	"errors"

	"github.com/hossainabid/go-ims/domain"
	"github.com/hossainabid/go-ims/models"
	"github.com/hossainabid/go-ims/types"
	"github.com/hossainabid/go-ims/utils/errutil"
)

type ProductServiceImpl struct {
	productRepo domain.ProductRepository
}

func NewProductServiceImpl(productRepo domain.ProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{
		productRepo: productRepo,
	}
}

func (svc *ProductServiceImpl) CreateProduct(productReq *types.CreateProductRequest) (*types.CreateProductResponse, error) {
	product := productReq.ToProduct()
	createdProduct, err := svc.productRepo.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return &types.CreateProductResponse{
		Message: "Product created",
		Product: createdProduct,
	}, nil
}

func (svc *ProductServiceImpl) ListProducts(productReq types.ListProductRequest) (*types.PaginatedProductResponse, error) {
	offset := (productReq.Page - 1) * productReq.Limit
	products, count, err := svc.productRepo.ListProducts(productReq.Limit, offset)
	if errors.Is(err, errutil.ErrRecordNotFound) {
		return &types.PaginatedProductResponse{}, nil
	}
	if err != nil {
		return nil, err
	}
	response := &types.PaginatedProductResponse{
		Page:     productReq.Page,
		Limit:    productReq.Limit,
		Total:    count,
		Products: products,
	}
	return response, nil
}

func (svc *ProductServiceImpl) ReadProductByID(id int) (*models.Product, error) {
	product, err := svc.productRepo.ReadProductByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (svc *ProductServiceImpl) UpdateProduct(productReq *types.UpdateProductRequest) (*types.UpdateProductResponse, error) {
	existingProduct, err := svc.productRepo.ReadProductByID(productReq.ID)
	if err != nil {
		return nil, err
	}
	if existingProduct == nil {
		return nil, errutil.ErrRecordNotFound
	}

	product := productReq.ToProduct()
	updatedProduct, err := svc.productRepo.UpdateProduct(product)
	if err != nil {
		return nil, err
	}
	return &types.UpdateProductResponse{
		Message: "Product updated",
		Product: updatedProduct,
	}, nil
}

func (svc *ProductServiceImpl) DeleteProduct(id int) (*types.DeleteProductResponse, error) {
	err := svc.productRepo.DeleteProduct(id)
	if err != nil {
		return nil, err
	}
	return &types.DeleteProductResponse{
		Message: "Product deleted",
	}, nil
}
