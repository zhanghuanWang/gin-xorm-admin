package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-contrib/sessions"
)

type AuthController struct {
}

func (AuthController) ToLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Admin",
	})
}

func (AuthController) Login(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("user_id", "1")
	session.Save()
	c.Redirect(http.StatusMovedPermanently, "/")
}

func (AuthController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusMovedPermanently, "/login")
}
