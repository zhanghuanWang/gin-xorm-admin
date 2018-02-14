package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"strings"
)

// ErrorHandler is a middleware to handle errors encountered during requests
func ErrorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		c.HTML(http.StatusBadRequest, "400", gin.H{
			"errors": c.Errors,
		})
	}
}

// NoRoute is a middleware to handle page not found during requests
func NoRoute(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404", gin.H{
		"title": "404",
	})
}

// Auth is a middleware to handle the authenticate
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL
		// 过滤资源文件，登录和退出
		if strings.HasPrefix(url.String(), "/public") || strings.Index(url.String(), "login") > -1 ||
			strings.Index(url.String(), "/logout") > -1 {
			c.Next()
			return
		}
		session := sessions.Default(c)
		id := session.Get("user_id")
		if id != nil {
			c.Next()
			return
		}
		c.Redirect(http.StatusFound, "login")
		c.Abort()
		return
	}
}
