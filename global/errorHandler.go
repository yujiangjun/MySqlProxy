package global

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"mySqlProxy/config"
	"net/http"
	"runtime/debug"
)

func Recover(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Error("panic:%v\n",r)
			debug.PrintStack()
			ctx.JSON(http.StatusOK,config.Error(errorToString(r)))
			ctx.Abort()
		}
	}()
	ctx.Next()
}

//errorToString recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}