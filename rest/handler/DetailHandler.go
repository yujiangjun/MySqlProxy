package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"mySqlProxy/config"
	"mySqlProxy/jdbc/dto"
	"net/http"
	"strconv"
	"strings"
)

func GetTableDesc(ctx *gin.Context) {
	databaseId, ok := ctx.GetQuery("databaseId")
	if !ok {
		log.Error("databaseId不能为空")
		ctx.JSON(http.StatusOK,config.Error("databaseId不能为空"))
	}
	schema, ok := ctx.GetQuery("schema")
	if !ok {
		log.Error("schema不能为空")
		ctx.JSON(http.StatusOK,config.Error("schema不能为空"))
	}
	table, ok := ctx.GetQuery("table")
	if !ok {
		log.Error("table不能为空")
		ctx.JSON(http.StatusOK,config.Error("table不能为空"))
		return
	}
	id, _ := strconv.Atoi(databaseId)
	db := GetConnectById(id)

	tableInfo := make(map[string]interface{})

	err := db.Raw("select * from information_schema.TABLES where TABLE_SCHEMA=? and TABLE_NAME=?", schema, table).Scan(&tableInfo)
	if err.Error != nil {
		log.Error("查询异常",err.Error)
	}
	ctx.JSON(http.StatusOK,config.Success(tableInfo))
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
		ctx.JSON(http.StatusOK,config.Error("查询异常"+result.Error.Error()))
		return
	}
	ctx.JSON(http.StatusOK,config.Success(columnInfo))
}

func AlertTab(ctx *gin.Context) {
	var req dto.AlertTab
	ctx.ShouldBindJSON(&req)
	addSql:=""
	for _, column := range req.Columns {
		addSql += fmt.Sprintf("`%s` %s default %s comment '%s', ", column.ColumnName, column.ColumnType, column.Default, column.Comment)
	}
	index := strings.LastIndex(addSql, ",")
	if index !=-1 {
		addSql=addSql[:index]
	}
	sql := fmt.Sprintf("alter table %s.%s add (%s)", req.Schema, req.Table,addSql)

	log.Info("sql:",sql)

	db := connectionMaps[req.DatabaseId]
	result := db.Exec(sql)
	if result.Error!=nil {
		log.Error("修改表发生错误:",result.Error)
		ctx.JSON(http.StatusOK,config.Error(fmt.Sprintf("修改表发生错误:%s",result.Error.Error())))
		return
	}
	ctx.JSON(http.StatusOK,config.Success(sql))
}

func DeleteCols(ctx *gin.Context) {
	var req dto.DeleteColForTabReq
	ctx.ShouldBindJSON(&req)
	log.Info("接收参数:",req)
}

func PingTest(ctx *gin.Context) {
	id, _ := ctx.GetQuery("id")
	if id == "1" {
		panic("id不能是1")
	}
}