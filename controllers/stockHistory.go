package controllers

import (
	"net/http"

	"github.com/hossainabid/go-ims/consts"
	"github.com/hossainabid/go-ims/middlewares"

	"github.com/hossainabid/go-ims/domain"
	"github.com/hossainabid/go-ims/logger"
	"github.com/hossainabid/go-ims/types"
	"github.com/hossainabid/go-ims/utils/msgutil"
	"github.com/labstack/echo/v4"
)

type StockHistoryController struct {
	stockHistorySvc domain.StockHistoryService
}

func NewStockHistoryController(stockHistorySvc domain.StockHistoryService) *StockHistoryController {
	return &StockHistoryController{
		stockHistorySvc: stockHistorySvc,
	}
}

func (ctrl *StockHistoryController) RecordStockHistory(c echo.Context) error {
	var req types.RecordStockHistoryRequest
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

	resp, err := ctrl.stockHistorySvc.RecordStockHistory(&req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, msgutil.SomethingWentWrongMsg())
	}
	return c.JSON(http.StatusCreated, resp)
}

func (ctrl *StockHistoryController) ListStockHistories(c echo.Context) error {
	req := types.ListStockHistoryRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, msgutil.InvalidRequestMsg())
	}

	if err := req.Validate(); err != nil {
		logger.Error("validation error: %v", err)
		return c.JSON(http.StatusBadRequest, &types.ValidationError{
			Error: err,
		})
	}

	if req.Limit <= 0 {
		req.Limit = consts.DefaultPageSize
	}
	if req.Page <= 0 {
		req.Page = consts.DefaultPage
	}
	stockHistories, err := ctrl.stockHistorySvc.ListStockHistories(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, msgutil.SomethingWentWrongMsg())
	}
	return c.JSON(http.StatusOK, stockHistories)
}
