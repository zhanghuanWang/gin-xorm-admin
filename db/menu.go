package db

import (
	"fmt"
	"github.com/angao/gin-xorm-admin/models"
)

// MenuDao 菜单操作
type MenuDao struct{}

// GetMenuByRoleIds 根据角色查询所属菜单
func (MenuDao) GetMenuByRoleIds(roleID int64) ([]models.Menu, error) {
	var menus []models.Menu
	err := X.SqlMapClient("GetMenuByRoleIds", roleID).Find(&menus)
	if err != nil {
		fmt.Printf("%#v\n", err)
		return nil, err
	}
	return menus, nil
}