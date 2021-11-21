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
		//后续修改。
		var result []map[string]interface{}
		exec := db.Raw(req.Sql).Scan(&result)
		if exec.Error!=nil {
			panic(exec.Error)
		}
		rows, _ := db.Raw(req.Sql).Rows()
		defer rows.Close()
		columns, _ := rows.Columns()
		types := make([]dto.ColumnType, len(columns))
		for i := 0; i < len(columns); i++ {
			types[i]=dto.ColumnType{
				Title:     columns[i],
				DataIndex: columns[i],
				Key:       columns[i],
			}
		}
		ctx.JSON(http.StatusOK,config.Success(struct {
			ColumnKey []dto.ColumnType             `json:"columnKey"`
			Result    []map[string]interface{} `json:"result"`
		}{
			ColumnKey: types,
			Result: result,
		}))
		return
	}
	exec := db.Exec(req.Sql)
	if exec.Error!=nil {
		panic(exec.Error)
	}
	ctx.JSON(http.StatusOK,config.Success(fmt.Sprintf("影响行数:%d",exec.RowsAffected)))
}
