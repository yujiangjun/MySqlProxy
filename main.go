package main

import (
	"github.com/gin-gonic/gin"
	"mySqlProxy/global"
	"net/http"
	"time"
)

func main() {
	router := InitRouter(gin.Default())
	global.NewGlobal()

	InitServer(":9999",router).ListenAndServe().Error()
}

func InitServer(address string, router *gin.Engine) *http.Server {
	gin.SetMode(gin.DebugMode)
	return &http.Server{
		Addr: address,
		Handler: router,
		ReadTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
		MaxHeaderBytes: 1<<20,
	}
}
