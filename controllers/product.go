package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/hossainabid/go-ims/consts"
	"github.com/hossainabid/go-ims/middlewares"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hossainabid/go-ims/domain"
	"github.com/hossainabid/go-ims/logger"
	"github.com/hossainabid/go-ims/types"
	"github.com/hossainabid/go-ims/utils/errutil"
	"github.com/hossainabid/go-ims/utils/msgutil"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productSvc domain.ProductService
}

func NewProductController(productSvc domain.ProductService) *ProductController {
	return &ProductController{
		productSvc: productSvc,
	}
}

func (ctrl *ProductController) CreateProduct(c echo.Context) error {
	var req types.CreateProductRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, msgutil.InvalidRequestMsg())
	}

	if err := req.Validate(); err != nil {
		logger.Error("validation error: %v", err)
		return c.JSON(http.StatusBadRequest, &types.ValidationError{
			Error: err,
		})
	}

	user, err := middlewares.CurrentUserFromCtx(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, msgutil.UserUnauthorized())
	}

	req.CreatedBy = user.ID

	resp, err := ctrl.productSvc.CreateProduct(&req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, msgutil.SomethingWentWrongMsg())
	}
	return c.JSON(http.StatusCreated, resp)
}

func (ctrl *ProductController) ListProducts(c echo.Context) error {
	req := types.ListProductRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, msgutil.InvalidRequestMsg())
	}
	if req.Limit <= 0 {
		req.Limit = consts.DefaultPageSize
	}
	if req.Page <= 0 {
		req.Page = consts.DefaultPage
	}
	products, err := ctrl.productSvc.ListProducts(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, msgutil.SomethingWentWrongMsg())
	}
	return c.JSON(http.StatusOK, products)
}

func (ctrl *ProductController) ReadProductByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, msgutil.InvalidRequestMsg())
	}

	// Validate ID
	if err := v.Validate(id, v.Required); err != nil {
		logger.Error("validation error: %v", err)
		return c.JSON(http.StatusBadRequest, &types.ValidationError{
			Error: err,
		})
	}

	product, err := ctrl.productSvc.ReadProductByID(id)
	if errors.Is(err, errutil.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, msgutil.ProductNotFound())
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, msgutil.SomethingWentWrongMsg())
	}
	return c.JSON(http.StatusOK, product)
}

func (ctrl *ProductController) UpdateProduct(c echo.Context) error {
	var req types.UpdateProductRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, msgutil.InvalidRequestMsg())
	}

	if err := req.Validate(); err != nil {
		logger.Error("validation error: %v", err)
		return c.JSON(http.StatusBadRequest, &types.ValidationError{
			Error: err,
		})
	}

	resp, err := ctrl.productSvc.UpdateProduct(&req)
	if errors.Is(err, errutil.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, msgutil.ProductNotFound())
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, msgutil.SomethingWentWrongMsg())
	}
	return c.JSON(http.StatusOK, resp)
}

func (ctrl *ProductController) DeleteProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, msgutil.InvalidRequestMsg())
	}

	// Validate ID
	if err := v.Validate(id, v.Required); err != nil {
		logger.Error("validation error: %v", err)
		return c.JSON(http.StatusBadRequest, &types.ValidationError{
			Error: err,
		})
	}

	resp, err := ctrl.productSvc.DeleteProduct(id)
	if errors.Is(err, errutil.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, msgutil.ProductNotFound())
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, msgutil.SomethingWentWrongMsg())
	}
	return c.JSON(http.StatusNoContent, resp)
}
