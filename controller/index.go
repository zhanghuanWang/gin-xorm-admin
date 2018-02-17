package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-contrib/sessions"
)

type IndexController struct {
}

func (IndexController) Home(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username").(string)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"username": username,
	})
}

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
