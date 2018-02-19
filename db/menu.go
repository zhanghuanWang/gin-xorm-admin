package db

import (
	"strconv"
	"github.com/angao/gin-xorm-admin/models"
)

type MenuDao struct{}

// GetMenuByRoleIds 根据角色查询所属菜单
func (MenuDao) GetMenuByRoleIds(roleID int64) ([]models.Menu, error) {
	sql := `
		SELECT
		m1.id AS id ,
		m1.icon AS icon ,
		(
			CASE
			WHEN(m2.id = 0 OR m2.id IS NULL) THEN
				0
			ELSE
				m2.id
			END
		) AS parentId ,
		m1.name AS name,
		m1.url AS url,
		m1.levels AS levels ,
		m1.ismenu AS ismenu ,
		m1.num AS num
	FROM
		sys_menu m1
	LEFT JOIN sys_menu m2 ON m1.pcode = m2.code
	INNER JOIN(
		SELECT
			ID
		FROM
			sys_menu
		WHERE
			ID IN(
				SELECT
					menuid
				FROM
					sys_relation rela
				WHERE
					rela.roleid = ?
			)
	) m3 ON m1.id = m3.id
	WHERE
		m1.ismenu = 1
	ORDER BY
		levels ,
		num ASC	
	`
	results, err := X.SQL(sql, roleID).QueryInterface()
	if err != nil {
		return nil, err
	}
	var menus []models.Menu
	for _, result := range results {
		var menu models.Menu
		if id, ok := result["id"].(int64); ok {
			menu.Id = id
		}
		if parentID, ok := result["parentId"].(int64); ok {
			menu.Pcode = strconv.Itoa(int(parentID))
		}
		if icon, ok := result["icon"].([]uint8); ok {
			menu.Icon = string(icon)
		}
		if name, ok := result["name"].([]uint8); ok {
			menu.Name = string(name)
		}
		if url, ok := result["url"].([]uint8); ok {
			menu.URL = string(url)
		}
		if levels, ok := result["levels"].(int64); ok {
			menu.Levels = int(levels)
		}
		if ismenu, ok := result["ismenu"].(int64); ok {
			menu.IsMenu = int(ismenu)
		}
		if num, ok := result["num"].(int64); ok {
			menu.Num = int(num)
		}
		menus = append(menus, menu)
	}
	return menus, nil
}