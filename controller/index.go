package controller

import (
	"github.com/angao/gin-xorm-admin/models"
	"fmt"
	"strconv"
	"github.com/angao/gin-xorm-admin/db"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// IndexController is handle home page request
type IndexController struct {
}

// Home is handle "/" request
func (IndexController) Home(c *gin.Context) {
	session := sessions.Default(c)
	userID, ok := session.Get("user_id").(int64)
	var user *models.UserRole
	var err error
	if ok {
		var userDao db.UserDao
		var menuDao db.MenuDao
		user, err = userDao.GetUserRole(userID)
		if err != nil {
			fmt.Printf("%#v\n", err)
		} else {
			roleID, _ := strconv.ParseInt(user.RoleId, 10, 64)
			_, err := menuDao.GetMenuByRoleIds(roleID)
			if err != nil {
				fmt.Printf("%#v\n", err)
				return
			}
			c.HTML(http.StatusOK, "index.html", gin.H{
				"username": user.User.Name,
				"rolename": user.Role.Name,
			})
		}
	}
} 

// BlackBoard is handle "/blackboard"
func (IndexController) BlackBoard(c *gin.Context) {
	notices := make([]map[string]string, 0)
	notice := map[string]string{
		"content": "欢迎使用Admin系统",
	}
	notices = append(notices, notice)
	c.HTML(http.StatusOK, "container.html", gin.H{
		"noticeList": notices,
	})
}
