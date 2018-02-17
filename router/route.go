package router

import (
	"html/template"
	"github.com/angao/gin-xorm-admin/router/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/angao/gin-xorm-admin/controller"
	"github.com/gin-contrib/sessions"
	"github.com/angao/gin-xorm-admin/utils"
)

// Init 路由
func Init() {
	router := gin.New()

	store := sessions.NewCookieStore([]byte("--secret--key--"))
	router.Use(sessions.Sessions("session_id", store))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.ErrorHandler)

	router.NoRoute(middlewares.NoRoute)

	router.Static("/public", "public")
	router.HTMLRender = utils.LoadTemplates("views")

	router.SetFuncMap(template.FuncMap{
		"formatAsDate": utils.FormatDate,
	})

	// login authentication
	auth := new(controller.AuthController)
	router.GET("/login", auth.ToLogin)
	router.POST("/login", auth.Login)
	router.GET("/logout", auth.Logout)

	index := new(controller.IndexController)
	router.GET("/", middlewares.Auth(), index.Home)
	router.GET("/blackboard", middlewares.Auth(), index.BlackBoard)

	// user
	user := new(controller.UserController)
	userGroup := router.Group("/user", middlewares.Auth())
	{
		userGroup.GET("/info", user.Info)
	}

	router.Run(":3000")
}
