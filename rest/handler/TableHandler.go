package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"mySqlProxy/config"
	"mySqlProxy/const"
	"mySqlProxy/global"
	"mySqlProxy/jdbc"
	"mySqlProxy/jdbc/dto"
	"mySqlProxy/jdbc/entity"
	"net/http"
	"strconv"
	"strings"
)

var connectionMaps = make(map[int]*gorm.DB)

func GetConnectById(id int) *gorm.DB {
	db := connectionMaps[id]
	return db
}

func GetFields(context *gin.Context) {
	databaseId, ok := context.GetQuery("databaseId")
	if !ok {
		log.Info("未指定databaseId")
		context.JSON(http.StatusOK,config.Error("未指定databaseId"))
		return
	}
	schema, ok := context.GetQuery("schema")
	if !ok {
		log.Info("未指定schema")
		context.JSON(http.StatusOK,config.Error("未指定schema"))
		return
	}
	table, ok := context.GetQuery("table")
	if !ok {
		log.Info("未指定table")
		context.JSON(http.StatusOK,config.Error("未指定table"))
		return
	}
	log.Info("databaseId:%s", databaseId)
	id, _ := strconv.Atoi(databaseId)
	log.Info(connectionMaps[id])
	var fields []map[string]interface{}
	err := connectionMaps[id].Raw("SELECT COLUMN_NAME fName,column_comment fDesc,DATA_TYPE dataType, IS_NULLABLE isNull,IFNULL(CHARACTER_MAXIMUM_LENGTH,0) sLength FROM information_schema.columns where TABLE_SCHEMA=? and TABLE_NAME=?",schema,table).Scan(&fields)
	if err.Error != nil {
		log.Error("查询失败", err.Error)
	}
	context.JSON(http.StatusOK, config.Success(fields))
}

func GetTables(ctx *gin.Context) {
	databaseId, ok := ctx.GetQuery("databaseId")
	if !ok {
		log.Error("未指定数据库")
		ctx.JSON(http.StatusOK,config.Error("未指定数据库"))
		return
	}
	schema, ok := ctx.GetQuery("schema")

	if !ok {
		log.Error("未指定schema")
		ctx.JSON(http.StatusOK,config.Error("未指定schema"))
		return
	}

	id, _ := strconv.Atoi(databaseId)

	conn := connectionMaps[id]
	log.Info("链接信息", conn)
	var tables []map[string]interface{}
	result := conn.Raw("select TABLE_CATALOG,TABLE_SCHEMA,TABLE_NAME,TABLE_TYPE,ENGINE,VERSION,ROW_FORMAT,TABLE_ROWS,DATA_LENGTH,CREATE_TIME,UPDATE_TIME,TABLE_COLLATION,TABLE_COMMENT from information_schema.TABLES where TABLE_SCHEMA=?",schema).Scan(&tables)
	if result.Error != nil {
		log.Error("查询发生异常.", result.Error)
	}
	ctx.JSON(http.StatusOK,config.Success(tables))
}

func GetSchemas(ctx *gin.Context) {
	databaseId, ok := ctx.GetQuery("databaseId")
	if !ok {
		log.Error("databaseId 不能为空")
		ctx.JSON(http.StatusOK,config.Error("databaseId 不能为空"))
		return
	}
	id, _ := strconv.Atoi(databaseId)
	db := connectionMaps[id]
	var schemas []map[string]interface{}
	err := db.Raw("select CATALOG_NAME,SCHEMA_NAME,DEFAULT_CHARACTER_SET_NAME,DEFAULT_COLLATION_NAME from information_schema.SCHEMATA").Scan(&schemas)
	if err!=nil {
		log.Error("查询发生异常.cause:",err)
	}
	ctx.JSON(http.StatusOK, config.Success(schemas))
}

func GetSchema(ctx *gin.Context) {
	databaseId, ok := ctx.GetQuery("databaseId")
	if !ok {
		log.Error("databaseId 不能为空")
		ctx.JSON(http.StatusOK,config.Error("databaseId 不能为空"))
		return
	}
	schema, ok := ctx.GetQuery("schema")
	if !ok {
		log.Error("schema不能为空")
		ctx.JSON(http.StatusOK,config.Error("schema不能为空"))
	}

	id, _ := strconv.Atoi(databaseId)
	db := connectionMaps[id]
	schemaInfo:=make(map[string]interface{})
	err := db.Raw("select * from information_schema.SCHEMATA where SCHEMA_NAME=?", schema).Scan(&schemaInfo)
	if err!=nil {
		log.Error("查询发生异常")
	}
	ctx.JSON(http.StatusOK, config.Success(schemaInfo))
}





func Login(ctx *gin.Context) {
	userName := ctx.Query("userName")
	password := ctx.Query("password")
	result := make(map[string]string)
	if strings.EqualFold(userName, "admin") && strings.EqualFold(password, "123") {
		result["code"] = "200"
		result["msg"] = "成功"
		ctx.JSON(http.StatusOK, config.Success("登陆成功"))
		return
	}
	ctx.JSON(http.StatusOK, config.Error("登陆失败"))
}


func DataBasePing(ctx *gin.Context) {

	dataSource := dto.DataSource{}
	ctx.BindJSON(&dataSource)

	log.Info("参数:", dataSource)
	log.Info("拼接的url:", fmt.Sprintf("tcp(%s:%d)/%s?charset=utf8&parseTime=true", dataSource.Host, dataSource.Port, dataSource.Schema))
	meta := jdbc.MyJdbc{
		Url:      fmt.Sprintf("tcp(%s:%d)/%s?charset=utf8&parseTime=true", dataSource.Host, dataSource.Port, dataSource.Schema),
		Username: dataSource.UserName,
		Password: dataSource.Password,
	}
	connection := meta.GetConnection()
	db, _ := connection.DB()
	err := db.Ping()
	if err != nil {
		ctx.JSON(http.StatusOK, config.Error("连接失败"))
		return
	}
	ctx.JSON(http.StatusOK,config.Success("连接成功"))
}

func CreateConnect(ctx *gin.Context) {
	dataSource := dto.DataSource{}
	ctx.BindJSON(&dataSource)
	log.Info("dataSource:", dataSource)
	myJdbc := jdbc.MyJdbc{
		Username: dataSource.UserName,
		Password: dataSource.Password,
		Url:      fmt.Sprintf("tcp(%s:%d)/%s?charset=utf8&parseTime=true", dataSource.Host, dataSource.Port, dataSource.Schema),
	}
	connection := myJdbc.GetConnection()
	db, _ := connection.DB()
	err2 := db.Ping()
	if err2 != nil {
		ctx.JSON(http.StatusOK, config.Error("连接失败"))
		return
	}
	metaService := new(jdbc.Meta)

	log.Info("记录链接信息", metaService.GetMeta())
	connect := global.GetGlobal().SqlConnect
	dataConnect := entity.DataConnect{
		DbName:   "test",
		DbType:   0,
		Url:      myJdbc.Url,
		UserName: myJdbc.Username,
		Password: myJdbc.Password,
	}
	//临时指定表明，默认是蛇形复数形式如这里是data_connects.
	create := connect.Table("data_connect").Create(&dataConnect)
	//想map中添加connection
	connectionMaps[dataConnect.Id] = connection

	if create.Error != nil {
		log.Error("保存数据库信息失败：", create.Error)
		ctx.JSON(http.StatusOK, config.Error("保存数据库信息失败"))
		return
	}
	log.Info("保存数据库信息成功。")

	redisHelper := global.GetGlobal().RedisConnect
	json, _ := json.Marshal(myJdbc)

	result, err := redisHelper.Set(context.Background(), fmt.Sprintf("datasource_%d", dataConnect.Id), json, 0).Result()
	if err != nil {
		ctx.JSON(http.StatusOK, config.Error("缓存到数据失败"))
		return
	}
	log.Info("缓存成功", result)

	ctx.JSON(http.StatusOK,config.Success("创建数据库链接成功"))
}

func GetRedisCache(ctx *gin.Context) {
	redis := global.GetGlobal().RedisConnect

	result, err := redis.Get(context.Background(), "datasource_1").Result()
	if err != nil {
		log.Error("获取缓存数据失败", err)
	}
	ctx.JSON(http.StatusOK,config.Success(result))
}

func InitLoadingConnection2Redis() {
	db := global.GetGlobal().SqlConnect
	redisDb := global.GetGlobal().RedisConnect
	var dataConnect []*entity.DataConnect
	db.Table("data_connect").Find(&dataConnect)
	for _, value := range dataConnect {

		connectionMaps[value.Id] = jdbc.MyJdbc{
			Username: value.UserName,
			Password: value.Password,
			Url:      value.Url,
		}.GetConnection()
		json, _ := json.Marshal(value)
		log.Info("加载到redis中:",value)
		err := redisDb.Set(context.Background(), fmt.Sprintf("datasource_%d", value.Id), json, 0).Err()
		if err != nil {
			log.Error("存入缓存缓存错误", err)
		}
	}
	log.Info("加载链接到缓存成功")
}

func GetDbs(ctx *gin.Context) {
	redis := global.GetGlobal().RedisConnect
	result, _ := redis.Keys(context.Background(), "datasource*").Result()

	conns := make([]*entity.DataConnect, len(result))
	for i := 0; i < len(result); i++ {
		conn, _ := redis.Get(context.Background(), result[i]).Result()
		json.Unmarshal([]byte(conn), &conns[i])
	}
	ctx.JSON(http.StatusOK,config.Success(conns))
}

type CreateTabReq struct {
	DatabaseId int `json:"databaseId"`
	CreateSql string `json:"createSql"`
}
func CreateTable(ctx *gin.Context) {
	var req CreateTabReq
	ctx.ShouldBindJSON(&req)
	log.Info("请求参数：",req)

	db := connectionMaps[req.DatabaseId]
	result := db.Exec(req.CreateSql)
	if result.Error!=nil {
		log.Error("创建表发生错误:",result.Error)
		panic(result.Error)
		return
	}
	ctx.JSON(http.StatusOK,config.Success(fmt.Sprintf("success.影响行数:%d",result.RowsAffected)))
}

//InsertDataForTab 插入数据
func InsertDataForTab(ctx *gin.Context) {
	var req dto.InsertDataForTab
	ctx.ShouldBindJSON(&req)
	log.Info("参数:",req)
	db := connectionMaps[req.DatabaseId]
	result:= db.Exec(req.InsertSql)
	if result.Error!=nil {
		log.Error("插入数据发生错误:",result.Error)
		ctx.JSON(http.StatusOK,config.Error("插入数据发生错误:"+result.Error.Error()))
		return
	}
	ctx.JSON(http.StatusOK,config.Success("插入数据发生错误:"+fmt.Sprintf("success.影响行数:%d",result.RowsAffected)))
}

//ChangeTabName 修改表名
func ChangeTabName(ctx *gin.Context) {
	var req dto.ChangeTabNameReq
	ctx.ShouldBindJSON(&req)
	log.Info("接受参数:",req)
	db := connectionMaps[req.DatabaseId]
	sql := fmt.Sprintf("ALTER TABLE %s.%s RENAME AS %s", req.Schema, req.OriTab, req.NewTab)
	if err:=db.Exec(sql).Error;err!=nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK,config.Success(nil))
}

func DropTab(ctx *gin.Context) {
	var req dto.DropTabReq
	ctx.ShouldBindJSON(&req)
	log.Info("接受参数:",req)
	sql := fmt.Sprintf("drop table %s.%s", req.Schema, req.Table)
	db := connectionMaps[req.DatabaseId]
	if result := db.Exec(sql); result.Error != nil {
		panic(result.Error)
	}
	ctx.JSON(http.StatusOK,config.Success(nil))
}

func GetDataTypeList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK,config.Success(_const.DataTypeList()))
}

func GetFullDataTab(ctx *gin.Context) {
	databaseId, ok := ctx.GetQuery("databaseId")
	if !ok {
		panic("缺少参数databaseId")
	}
	schema, ok := ctx.GetQuery("schema")
	if !ok {
		panic("缺少参数schema");
	}
	table, ok := ctx.GetQuery("table")
	if !ok {
		panic("缺少参数table")
	}
	sql := fmt.Sprintf("select * from %s.%s", schema, table)
	id, _ := strconv.Atoi(databaseId)

	db := connectionMaps[id]

	var results []map[string]interface{}

	exec := db.Raw(sql).Scan(&results)
	if exec.Error!=nil {
		panic(exec.Error)
	}

	colInfoSql:="select * from information_schema.COLUMNS where TABLE_SCHEMA=? and TABLE_NAME=?"
	var columnInfo  []map[string]interface{}

	exec = db.Raw(colInfoSql, schema,table).Scan(&columnInfo)

	if exec.Error!=nil {
		panic(exec.Error)
	}
	//var columnKeys [len(columnInfo)]ColumnType
	columnKeys := make([]dto.ColumnType, len(columnInfo))
	for i := 0; i < len(columnInfo); i++ {
		columnKeys[i]=dto.ColumnType{
			Title:     columnInfo[i]["COLUMN_NAME"].(string),
			DataIndex: columnInfo[i]["COLUMN_NAME"].(string),
			Key:       columnInfo[i]["COLUMN_NAME"].(string),
		}
	}
	ctx.JSON(http.StatusOK,config.Success(struct {
		ColumnKey []dto.ColumnType             `json:"columnKey"`
		Result    []map[string]interface{} `json:"result"`
	}{
		ColumnKey: columnKeys,
		Result: results,
	}))
}