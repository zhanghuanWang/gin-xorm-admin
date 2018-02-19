package db

import (
	"github.com/angao/gin-xorm-admin/models"
)

type NoticeDao struct{}

// List query all notice
func (NoticeDao) List() ([]models.Notice, error) {
	var notices []models.Notice
	err := X.Table("sys_notice").Find(&notices)
	if err != nil {
		return nil, err
	}
	return notices, nil
}