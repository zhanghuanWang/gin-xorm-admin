package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"log"
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
		session := sessions.Default(c)
		id := session.Get("user_id").(int64)
		log.Printf("id: %#v\n", id)
		if id > 0 {
			c.Next()
			return
		}
		c.Redirect(http.StatusFound, "login")
		c.Abort()
		return
	}
}
