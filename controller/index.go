package controller

import (
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
	var username string
	var userDao db.UserDao
	if ok {
		user, err := userDao.Get(userID)
		if err != nil {
			
		} else {
			username = user.Name
		}
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"username": username,
	})
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
