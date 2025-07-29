package controllers

import (
	"errors"
	"net/http"

	"github.com/hossainabid/go-ims/domain"
	"github.com/hossainabid/go-ims/middlewares"
	"github.com/hossainabid/go-ims/types"
	"github.com/hossainabid/go-ims/utils/errutil"
	"github.com/hossainabid/go-ims/utils/msgutil"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authSvc domain.AuthService
}

func NewAuthController(authSvc domain.AuthService) *AuthController {
	return &AuthController{authSvc: authSvc}
}

func (ctrl *AuthController) Login(c echo.Context) error {
	var req types.LoginReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, msgutil.InvalidRequestMsg())
	}

	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, &types.ValidationError{
			Error: err,
		})
	}

	resp, err := ctrl.authSvc.Login(&req)
	if err != nil {
		switch {
		case errors.Is(err, errutil.ErrUserNotFound):
			return c.JSON(http.StatusUnauthorized, msgutil.InvalidLoginCredentials())
		case errors.Is(err, errutil.ErrInvalidLoginCredentials):
			return c.JSON(http.StatusUnauthorized, msgutil.InvalidLoginCredentials())
		}
		return c.JSON(http.StatusInternalServerError, &types.CommonError{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctrl *AuthController) Logout(c echo.Context) error {
	user, err := middlewares.CurrentUserFromCtx(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, msgutil.UserUnauthorized())
	}

	if err := ctrl.authSvc.Logout(user.AccessUuid, user.RefreshUuid); err != nil {
		return c.JSON(http.StatusInternalServerError, &types.CommonError{
			Error: err.Error(),
		})
	}

	return c.NoContent(http.StatusOK)
}
