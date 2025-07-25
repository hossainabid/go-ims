package middlewares

import (
	"fmt"

	"github.com/hossainabid/go-ims/types"
	"github.com/labstack/echo/v4"
)

func CurrentUserFromCtx(c echo.Context) (*types.CurrentUser, error) {
	user, ok := c.Get(ContextKeyCurrentUser).(types.CurrentUser)
	if !ok {
		return nil, fmt.Errorf("user not found in request")
	}
	return &user, nil
}
