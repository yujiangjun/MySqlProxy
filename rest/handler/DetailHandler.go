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

//AlertTab 添加字段
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
	deleteSql:=""
	for _, columnName := range *req.Columns {
		deleteSql+=fmt.Sprintf("DROP %s,",columnName)
	}
	sql:=fmt.Sprintf("alter table %s.%s %s",*req.Schema,*req.Table,deleteSql)
	db := connectionMaps[*req.DatabaseId]
	sql=sql[:strings.LastIndex(sql,",")]
	result := db.Exec(sql)
	if result.Error!=nil {
		panic(result.Error)
		return
	}
	ctx.JSON(http.StatusOK,config.Success(nil))
}

//ChangeTabCols 更改字段
func ChangeTabCols(ctx *gin.Context) {
	var req dto.ChangeColReq
	ctx.ShouldBindJSON(&req)
	log.Info("接受参数:",req)

	db := connectionMaps[req.DatabaseId]

	for _, column := range req.Columns {
		sql := fmt.Sprintf("alter table %s.%s change %s %s %s", req.Schema, req.Table, column.OriColName, column.NewColName, column.ColumnType)
		log.Info("sql:",sql)
		if err:=db.Exec(sql).Error;err!=nil {
			panic(err)
		}
	}
	ctx.JSON(http.StatusOK,config.Success(nil))
}