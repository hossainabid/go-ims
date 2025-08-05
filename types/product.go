package types

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hossainabid/go-ims/models"
)

type (
	CreateProductRequest struct {
		Name         string  `json:"name"`
		Description  *string `json:"description"`
		Sku          string  `json:"sku"`
		ThresholdQty int     `json:"threshold_qty"`
		CreatedBy    int     `json:"created_by"`
	}

	UpdateProductRequest struct {
		ID int `param:"id"`
		CreateProductRequest
	}

	CreateProductResponse struct {
		Message string          `json:"message"`
		Product *models.Product `json:"product"`
	}

	DeleteProductResponse struct {
		Message string `json:"message"`
	}

	UpdateProductResponse struct {
		Message string          `json:"message"`
		Product *models.Product `json:"product"`
	}

	ListProductRequest struct {
		Page  int `query:"page"`
		Limit int `query:"limit"`
	}

	PaginatedProductResponse struct {
		Total    int               `json:"total"`
		Page     int               `json:"page"`
		Limit    int               `json:"limit"`
		Products []*models.Product `json:"products"`
	}
)

func (cpreq *CreateProductRequest) Validate() error {
	return v.ValidateStruct(cpreq,
		v.Field(&cpreq.Name, v.Required),
		v.Field(&cpreq.Description, v.When(cpreq.Description != nil, v.Length(0, 500))),
		v.Field(&cpreq.Sku, v.Required),
	)
}

func (upreq *UpdateProductRequest) Validate() error {
	return v.ValidateStruct(upreq,
		v.Field(&upreq.ID, v.Required),
		v.Field(&upreq.CreateProductRequest, v.Required),
	)
}

func (cpreq *CreateProductRequest) ToProduct() *models.Product {
	product := &models.Product{
		Name:         cpreq.Name,
		Description:  cpreq.Description,
		Sku:          cpreq.Sku,
		ThresholdQty: cpreq.ThresholdQty,
		CreatedBy:    cpreq.CreatedBy,
	}
	return product
}

func (upreq *UpdateProductRequest) ToProduct() *models.Product {
	product := &models.Product{
		ID:           upreq.ID,
		Name:         upreq.Name,
		Description:  upreq.Description,
		Sku:          upreq.Sku,
		ThresholdQty: upreq.ThresholdQty,
		CreatedBy:    upreq.CreatedBy,
	}
	return product
}
