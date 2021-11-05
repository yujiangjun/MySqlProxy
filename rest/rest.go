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

func main() {
	var ctx=context.Background()
	rdb := config.NewRedisHelper()
	if _,err:=rdb.Ping(ctx).Result();err!=nil {
		log.Error("redis链接失败.",err.Error())
	}
	gin.SetMode(gin.DebugMode)
	router:=gin.Default()
	router.Use(config.Cors())
	router.Use(func(context *gin.Context) {
		log.Info("this is a middleware")
	})
	router.GET("/test", func(context *gin.Context) {
		value, exist := context.GetQuery("name")
		if !exist {
			value="this key is not exist"
		}
		context.Data(http.StatusOK,"text/plain",[]byte(fmt.Sprintf("get Success!%s\n",value)))
	})

	router.GET("/getTable",handler.GetContext)
	router.GET("/getTables",handler.GetTables)
	router.GET("/login",handler.Login)
	router.POST("/ping",handler.DataBasePing)
	router.POST("/createConn",handler.CreateConnect)
	router.GET("/getRedisCache",handler.GetRedisCache)
	err := http.ListenAndServe(":9999", router)
	if err != nil {
		log.Error("服务器发生错误",err)
		os.Exit(-1)
		return
	}
}