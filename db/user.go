package db

import (
	"github.com/angao/gin-xorm-admin/models"
	"errors"
	"log"
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
func (UserDao) Get(id int64) (*models.User, error) {
	user := new(models.User)
	user.Id = id
	has, err := X.Table("sys_user").Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("user not found")
	}
	log.Printf("user: %#v\n", user)
	return user, nil
}
