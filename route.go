package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
	"mySqlProxy/config"
	"mySqlProxy/rest/handler"
)

func InitRouter(engine *gin.Engine) *gin.Engine {
	gin.ForceConsoleColor()

	gin.DefaultWriter= colorable.NewColorableStdout()
	engine.Use(config.Cors())
	//engine.Use(logger.ToFile())
	engine.Use(func(context *gin.Context) {
		log.Info("this is a middleware")
	})
	group := engine.Group("/")
	group.GET("/getDbs", handler.GetDbs)
	group.GET("/getFields", handler.GetFields)
	group.GET("/getTables", handler.GetTables)
	group.GET("/getSchemas", handler.GetSchemas)
	group.GET("/getSchema",handler.GetSchema)
	group.GET("/login", handler.Login)
	group.POST("/ping", handler.DataBasePing)
	group.POST("/createConn", handler.CreateConnect)
	group.GET("/getRedisCache", handler.GetRedisCache)
	group.GET("/getTableInfo",handler.GetTableDesc)
	group.GET("/getTableColumnInfo",handler.GetColumnInfo)
	group.POST("/createTab",handler.CreateTable)

	return engine
}
