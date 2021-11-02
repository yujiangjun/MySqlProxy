package config

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method:=context.Request.Method

		context.Header("Access-Control-Allow-Origin", "*")
		//服务器支持的所有跨域请求的方法
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		//允许跨域设置可以返回其他子段，可以自定义字段
		context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length,  contenttype , AccessToken,X-CSRF-Token,Token")
		// 允许浏览器（客户端）可以解析的头部 （重要）
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Content-Type")
		//设置缓存时间
		context.Header("Access-Control-Max-Age", "172800")
		//允许客户端传递校验信息比如 cookie (重要)
		context.Header("Access-Control-Allow-Credentials", "true")

		if method=="OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		//defer func() {
		//	if err := recover(); err != nil {
		//		log.Error("Panic info is:%v",err)
		//	}
		//}()
		context.Next()
	}
}
