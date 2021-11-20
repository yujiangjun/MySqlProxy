package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"mySqlProxy/config"
	"mySqlProxy/jdbc/dto"
	"net/http"
	"strings"
)

func ExecSql(ctx *gin.Context) {
	var req dto.ExecSql
	ctx.ShouldBindJSON(&req)
	log.Info("接受参数:",req)

	db := GetConnectById(req.DatabaseId)
	db.Exec(fmt.Sprintf("use %s",req.Schema))
	if strings.HasPrefix(req.Sql,"select") || strings.HasPrefix(req.Sql,"SELECT") {
		var result []map[string]interface{}
		exec := db.Raw(req.Sql).Scan(&result)
		if exec.Error!=nil {
			panic(exec.Error)
		}
		ctx.JSON(http.StatusOK,config.Success(result))
		return
	}
	exec := db.Exec(req.Sql)
	if exec.Error!=nil {
		panic(exec.Error)
	}
	ctx.JSON(http.StatusOK,config.Success(fmt.Sprintf("影响行数:%d",exec.RowsAffected)))
}
