package router

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/angao/gin-xorm-admin/router/middlewares"
	"github.com/angao/gin-xorm-admin/router/multitemplate"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"github.com/angao/gin-xorm-admin/controller"
	"github.com/gin-contrib/sessions"
)

// Init 路由
func Init() {
	router := gin.New()

	store := sessions.NewCookieStore([]byte("--secret--key--"))
	router.Use(sessions.Sessions("session_id", store))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.Auth())
	router.Use(middlewares.ErrorHandler)

	router.NoRoute(middlewares.NoRoute)

	router.Static("/public", "public")
	router.HTMLRender = loadTemplates("views")

	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatDate,
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/blackboard", func(c *gin.Context) {
		notices := make([]map[string]string, 0)
		notice := map[string]string{
			"content": "欢迎使用Admin系统",
		}
		notices = append(notices, notice)
		c.HTML(http.StatusOK, "container.html", gin.H{
			"noticeList": notices,
		})
	})
	// 登录认证
	auth := new(controller.AuthController)
	router.GET("/login", auth.ToLogin)
	router.POST("/login", auth.Login)
	router.GET("/logout", auth.Logout)

	// User
	user := new(controller.UserController)
	userGroup := router.Group("/user")
	{
		userGroup.GET("/info", user.Info)
	}

	router.Run(":3000")
}

func formatDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func loadTemplates(templatesDir string) multitemplate.Render {
	r := multitemplate.New()

	layouts, err := filepath.Glob(templatesDir + "/*.html")
	if err != nil {
		panic(err.Error())
	}

	commons, err := filepath.Glob(templatesDir + "/common/*.html")
	if err != nil {
		panic(err.Error())
	}

	systems, err := filepath.Glob(templatesDir + "/system/**/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, layout := range layouts {
		files := append([]string{layout}, systems...)
		files = append(files, commons...)
		r.Add(filepath.Base(layout), template.Must(template.ParseFiles(files...)))
	}
	return r
}
