package domain

import (
	"github.com/hossainabid/go-ims/models"
	"github.com/hossainabid/go-ims/types"
)

type (
	UserService interface {
		CreateUser(req *types.CreateUserReq) error
		UpdateUser(req *types.UpdateUserReq) error
		ReadUser(id int, fromCache bool) (*types.UserInfo, error)
		DeleteUser(id int) error
		ReadUserByEmail(email string) (*models.User, error)
		ListUsers(req types.ListUserReq) (*types.PaginatedUserResp, error)
		ReadPermissionsByRole(roleID int) ([]*models.Permission, error)
		StoreInCache(user *types.UserInfo) error
	}

	UserRepository interface {
		CreateUser(user *models.User) (*models.User, error)
		UpdateUser(user *models.User) error
		ReadUser(id int) (*models.User, error)
		DeleteUser(id int) error
		ReadUserByEmail(email string) (*models.User, error)
		ListUsers(limit, offset int) ([]*types.UserInfo, int, error)
		ReadPermissionsByRole(roleID int) ([]*models.Permission, error)
		UserCountByEmail(email string) (int, error)
	}
)
