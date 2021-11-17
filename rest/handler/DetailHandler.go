package handler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func GetTableDesc(ctx *gin.Context) {
	databaseId, ok := ctx.GetQuery("databaseId")
	if !ok {
		log.Error("databaseId不能为空")
		ctx.JSON(http.StatusOK,gin.H{
			"code":500,
			"msg":"databaseId不能为空",
		})
	}
	schema, ok := ctx.GetQuery("schema")
	if !ok {
		log.Error("schema不能为空")
		ctx.JSON(http.StatusOK,gin.H{
			"code":500,
			"msg":"schema不能为空",
		})
	}
	table, ok := ctx.GetQuery("table")
	if !ok {
		log.Error("stable不能为空")
	}
	id, _ := strconv.Atoi(databaseId)
	db := GetConnectById(id)

	tableInfo := make(map[string]interface{})

	err := db.Raw("select * from information_schema.TABLES where TABLE_SCHEMA=? and TABLE_NAME=?", schema, table).Scan(&tableInfo)
	if err.Error != nil {
		log.Error("查询异常",err.Error)
	}
	ctx.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"success",
		"data":tableInfo,
	})
}

type GetInfoReq struct {
	DatabaseId int16  `json:"databaseId" uri:"databaseId" form:"databaseId"`
	Schema     string `json:"schema" uri:"schema" form:"schema"`
	Table      string `json:"table" uri:"table" form:"table"`
	Column     string `json:"column" uri:"column" form:"column"`
}
func GetColumnInfo(ctx *gin.Context) {
	var req GetInfoReq
	ctx.ShouldBindQuery(&req)
	log.Info("请求:",req)
	sql:="select * from information_schema.COLUMNS where TABLE_SCHEMA=? and TABLE_NAME=? and COLUMN_NAME=?"
	columnInfo := make(map[string]interface{})
	db := GetConnectById(int(req.DatabaseId))
	result := db.Raw(sql, req.Schema, req.Table, req.Column).Scan(&columnInfo)
	if  result.Error!=nil{
		log.Error("查询发生异常",result.Error)
		ctx.JSON(http.StatusOK,gin.H{
			"code":500,
			"msg":"查询异常"+result.Error.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"success",
		"data":columnInfo,
	})
}