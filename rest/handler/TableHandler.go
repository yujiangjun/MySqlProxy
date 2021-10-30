package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"mySqlProxy/jdbc"
	"net/http"
)

func GetContext( context *gin.Context){
	tableName, ok := context.GetQuery("tableName")
	if !ok  {
		log.Info("缺少tableName")
		return
	}
	log.Info("tableName:%s",tableName)
	metaService := new(jdbc.Meta)
	meta:=metaService.GetMeta()
	connection:= jdbc.GetConnection(meta)
	log.Info(connection)
	var tables []jdbc.Field
	err := connection.Select(&tables, "SELECT COLUMN_NAME fName,column_comment fDesc,DATA_TYPE dataType, IS_NULLABLE isNull,IFNULL(CHARACTER_MAXIMUM_LENGTH,0) sLength FROM information_schema.columns where TABLE_SCHEMA='raytine' and TABLE_NAME='apply_record'")
	if err!=nil {
		log.Error("查询失败",err)
	}
	log.Info("select success",tables)
	for _,value:=range tables{
		log.Info(value)
	}
	context.Data(http.StatusOK,"text/plain",[]byte(fmt.Sprintf("get Success!")))
}

func GetTables(ctx *gin.Context) {
	schema, ok := ctx.GetQuery("schema")
	if !ok {
		log.Error("为获取到参数schema")
		return
	}
	metaService := new(jdbc.Meta)
	meta := metaService.GetMeta()
	db := jdbc.GetConnection(meta)
	var tables []jdbc.Tables
	err := db.Select(&tables, "select TABLE_CATALOG,TABLE_SCHEMA,TABLE_NAME,TABLE_TYPE,ENGINE,VERSION,ROW_FORMAT,TABLE_ROWS,DATA_LENGTH,CREATE_TIME,UPDATE_TIME,TABLE_COLLATION,TABLE_COMMENT from information_schema.TABLES where TABLE_SCHEMA=?", schema)
	if err!=nil {
		log.Error("查询发生异常.",err)
	}
	for _,value:=range tables{
		log.Info(value)
	}
	ctx.JSON(http.StatusOK,tables)
}