package db

import (
	"fmt"
	"log"

	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/gin-gonic/gin"
)

// X 全局DB
var X *xorm.Engine

func init() {
	var err error
	cfg, err := ini.Load("config/db.ini")
	if err != nil {
		log.Fatal(err)
	}

	username := cfg.Section("mysql").Key("username").Value()
	password := cfg.Section("mysql").Key("password").Value()
	url := cfg.Section("mysql").Key("url").Value()

	source := fmt.Sprintf("%s:%s@%s", username, password, url)
	X, err = xorm.NewEngine("mysql", source)

	if err != nil {
		log.Fatal(err)
	}
	tablePrefix := core.NewPrefixMapper(core.SnakeMapper{}, "sys_")
	X.SetTableMapper(tablePrefix)
	X.Ping()
	X.ShowSQL(gin.IsDebugging())
}
