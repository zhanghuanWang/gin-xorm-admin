package db

import (
	"github.com/angao/gin-xorm-admin/models"
	"errors"
)

// UserDao operate user
type UserDao struct {
}

// GetUser query user by account
func (UserDao) GetUser(account string) (*models.User, error) {
	user := new(models.User)
	has, err := X.Table("sys_user").Where("account = ?", account).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// Get query user by primary key
func (UserDao) GetUserRole(id int64) (*models.UserRole, error) {
	user := new(models.UserRole)
	has, err := X.Table("sys_user").Join("INNER", "sys_role", "sys_user.roleid = sys_role.id").Where("sys_user.id = ?", id).Get(user)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("user not found")
	}
	return user, nil
}
