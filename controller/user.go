package controller

import (
	"log"
	"net/http"
	"github.com/angao/gin-xorm-admin/models"
	"github.com/gin-contrib/sessions"
	"github.com/angao/gin-xorm-admin/db"
	"github.com/gin-gonic/gin"
)

// UserController handle user request
type UserController struct {
}

// Info is handle user info
func (UserController) Info(c *gin.Context) {
	var userDao db.UserDao
	var err error
	var user *models.UserRole
	session := sessions.Default(c)
	id, ok := session.Get("user_id").(int64) 
	if ok {
		user, err = userDao.GetUserRole(id)
		if err != nil {
			log.Printf("%#v\n", err)
			return 
		}
		c.HTML(http.StatusOK, "container.html", gin.H{
			"user": user.User,
			"roleName": user.Role.Name,
		})
		return
	}
	log.Printf("id = %#v\n", id)
	c.HTML(http.StatusInternalServerError, "container.html", gin.H{
		"error": err,
		"user": user,
	})
}
