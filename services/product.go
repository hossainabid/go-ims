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
	userRepo    domain.UserRepository
}

func NewProductServiceImpl(productRepo domain.ProductRepository, userRepo domain.UserRepository) *ProductServiceImpl {
	return &ProductServiceImpl{
		productRepo: productRepo,
		userRepo:    userRepo,
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

func (svc *ProductServiceImpl) ListProducts(req types.ListProductRequest, user *types.CurrentUser) (*types.PaginatedProductResponse, error) {
	offset := (req.Page - 1) * req.Limit
	filter := svc.getProductListFilter(user)
	products, count, err := svc.productRepo.ListProducts(filter, req.Limit, offset)
	if errors.Is(err, errutil.ErrRecordNotFound) {
		return &types.PaginatedProductResponse{}, nil
	}
	if err != nil {
		return nil, err
	}
	response := &types.PaginatedProductResponse{
		Page:     req.Page,
		Limit:    req.Limit,
		Total:    count,
		Products: products,
	}
	return response, nil
}

func (svc *ProductServiceImpl) getProductListFilter(user *types.CurrentUser) *types.ProductFilter {
	filter := &types.ProductFilter{}
	return filter
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
