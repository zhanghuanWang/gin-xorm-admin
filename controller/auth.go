package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"

	"github.com/angao/gin-xorm-admin/db"
	"github.com/angao/gin-xorm-admin/utils"
)

// AuthController handle auth request
type AuthController struct {
}

// ToLogin to login page
func (AuthController) ToLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

// Login handle login 
func (AuthController) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"tips": "用户名密码不能为空",
		})
		return
	}

	var userDao db.UserDao
	user, err := userDao.GetUser(username)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"tips": err.Error(),
		})
		return
	}
	passwd, err := utils.Encrypt(password, user.Salt)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"tips": err.Error(),
		})
		return
	}
	if user.Password == passwd {
		session := sessions.Default(c)
		session.Set("user_id", user.Id)
		session.Save()
		c.Redirect(http.StatusMovedPermanently, "/")
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"tips": "密码错误",
		})
	}
}

// Logout is log out system
func (AuthController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusMovedPermanently, "/login")
}
