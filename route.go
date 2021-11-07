package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"mySqlProxy/config"
	"mySqlProxy/rest/handler"
	"net/http"
	"os"
)

func main1() {
	var ctx=context.Background()
	rdb := config.NewRedisHelper()
	if _,err:=rdb.Ping(ctx).Result();err!=nil {
		log.Error("redis链接失败.",err.Error())
	}
	gin.SetMode(gin.DebugMode)
	router:=gin.Default()
	router.GET("/test", func(context *gin.Context) {
		value, exist := context.GetQuery("name")
		if !exist {
			value="this key is not exist"
		}
		context.Data(http.StatusOK,"text/plain",[]byte(fmt.Sprintf("get Success!%s\n",value)))
	})


	err := http.ListenAndServe(":9999", router)
	if err != nil {
		log.Error("服务器发生错误",err)
		os.Exit(-1)
		return
	}
}

func InitRouter(engine *gin.Engine) *gin.Engine {
	engine.Use(config.Cors())
	engine.Use(func(context *gin.Context) {
		log.Info("this is a middleware")
	})
	group := engine.Group("/")
	group.GET("/getTable",handler.GetContext)
	group.GET("/getTables",handler.GetTables)
	group.GET("/login",handler.Login)
	group.POST("/ping",handler.DataBasePing)
	group.POST("/createConn",handler.CreateConnect)
	group.GET("/getRedisCache",handler.GetRedisCache)

	return engine
}