package db

import (
	"github.com/hossainabid/go-ims/models"
	"github.com/hossainabid/go-ims/types"
)

func (repo *Repository) CreateUser(user *models.User) (*models.User, error) {
	err := repo.client.Create(&user).Error
	return user, err
}

func (repo *Repository) ReadUser(id int) (*models.User, error) {
	var user models.User
	if err := repo.client.Model(&models.User{}).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *Repository) ListUsers(limit, offset int) ([]*types.UserInfo, int, error) {
	var users []*types.UserInfo
	var total int64

	if err := repo.client.Model(&models.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := repo.client.Model(&models.User{}).Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, int(total), nil
}

func (repo *Repository) ReadUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := repo.client.Model(&models.User{}).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *Repository) UserCountByEmail(email string) (int, error) {
	var total int64

	if err := repo.client.Model(&models.User{}).Where("email = ?", email).Count(&total).Error; err != nil {
		return 0, err
	}
	return int(total), nil
}

func (repo *Repository) ReadPermissionsByRole(roleID int) ([]*models.Permission, error) {
	var permissions []*models.Permission

	if err := repo.client.Model(&models.RolePermission{}).
		Select("permissions.*").
		Joins("JOIN permissions ON role_permissions.permission_id = permissions.id").
		Where("role_permissions.role_id = ?", roleID).Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}

func (repo *Repository) UpdateUser(user *models.User) error {
	updUserMap := map[string]interface{}{
		"first_name": user.FirstName,
		"last_name":  user.LastName,
	}
	return repo.client.Model(&models.User{}).
		Where("id = ?", user.ID).
		Updates(&updUserMap).Error
}

func (repo *Repository) DeleteUser(id int) error {
	if err := repo.client.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
