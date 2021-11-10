package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"mySqlProxy/config"
	"mySqlProxy/logger"
	"mySqlProxy/rest/handler"
)

func InitRouter(engine *gin.Engine) *gin.Engine {
	engine.Use(config.Cors())
	engine.Use(logger.ToFile())
	engine.Use(func(context *gin.Context) {
		log.Info("this is a middleware")
	})
	group := engine.Group("/")
	group.GET("/getDbs", handler.GetDbs)
	group.GET("/getTable", handler.GetContext)
	group.GET("/getTables", handler.GetTables)
	group.GET("/login", handler.Login)
	group.POST("/ping", handler.DataBasePing)
	group.POST("/createConn", handler.CreateConnect)
	group.GET("/getRedisCache", handler.GetRedisCache)

	return engine
}
