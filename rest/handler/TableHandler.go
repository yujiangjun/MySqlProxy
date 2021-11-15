package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"mySqlProxy/global"
	"mySqlProxy/jdbc"
	"mySqlProxy/jdbc/dto"
	"mySqlProxy/jdbc/entity"
	"net/http"
	"strconv"
	"strings"
)

var connectionMaps = make(map[int]*gorm.DB)

func GetFields(context *gin.Context) {
	databaseId, ok := context.GetQuery("databaseId")
	if !ok {
		log.Info("未指定databaseId")
		context.JSON(http.StatusOK,gin.H{
			"code":"500",
			"msg":"未指定databaseId",
		})
		return
	}
	schema, ok := context.GetQuery("schema")
	if !ok {
		log.Info("未指定schema")
		context.JSON(http.StatusOK,gin.H{
			"code":"500",
			"msg":"未指定schema",
		})
		return
	}
	table, ok := context.GetQuery("table")
	if !ok {
		log.Info("未指定table")
		context.JSON(http.StatusOK,gin.H{
			"code":"500",
			"msg":"未指定table",
		})
		return
	}
	log.Info("databaseId:%s", databaseId)
	id, _ := strconv.Atoi(databaseId)
	log.Info(connectionMaps[id])
	var fields []map[string]interface{}
	err := connectionMaps[id].Raw("SELECT COLUMN_NAME fName,column_comment fDesc,DATA_TYPE dataType, IS_NULLABLE isNull,IFNULL(CHARACTER_MAXIMUM_LENGTH,0) sLength FROM information_schema.columns where TABLE_SCHEMA=? and TABLE_NAME=?",schema,table).Scan(&fields)
	if err != nil {
		log.Error("查询失败", err.Error)
	}
	context.JSON(http.StatusOK, gin.H{
		"code":200,
		"msg":"success",
		"data":fields,
	})
}

func GetTables(ctx *gin.Context) {
	databaseId, ok := ctx.GetQuery("databaseId")
	if !ok {
		log.Error("未指定数据库")
		return
	}
	schema, ok := ctx.GetQuery("schema")

	if !ok {
		log.Error("未指定schema")
		return
	}

	id, _ := strconv.Atoi(databaseId)

	conn := connectionMaps[id]
	log.Info("链接信息", conn)
	var tables []map[string]interface{}
	err := conn.Raw("select TABLE_CATALOG,TABLE_SCHEMA,TABLE_NAME,TABLE_TYPE,ENGINE,VERSION,ROW_FORMAT,TABLE_ROWS,DATA_LENGTH,CREATE_TIME,UPDATE_TIME,TABLE_COLLATION,TABLE_COMMENT from information_schema.TABLES where TABLE_SCHEMA=?",schema).Scan(&tables)
	if err != nil {
		log.Error("查询发生异常.", err.Error)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":200,
		"msg":"success",
		"data":tables,
	})
}

func GetSchemas(ctx *gin.Context) {
	databaseId, ok := ctx.GetQuery("databaseId")
	if !ok {
		log.Error("databaseId 不能为空")
		ctx.JSON(http.StatusOK,gin.H{
			"code":500,
			"msg":"未指定databaseId",
		})
		return
	}
	id, _ := strconv.Atoi(databaseId)
	db := connectionMaps[id]
	var schemas []map[string]interface{}
	err := db.Raw("select CATALOG_NAME,SCHEMA_NAME,DEFAULT_CHARACTER_SET_NAME,DEFAULT_COLLATION_NAME from information_schema.SCHEMATA").Scan(&schemas)
	if err!=nil {
		log.Error("查询发生异常.cause:",err)
	}
	ctx.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"success",
		"data":schemas,
	})
}

func GetSchema(ctx *gin.Context) {
	databaseId, ok := ctx.GetQuery("databaseId")
	if !ok {
		log.Error("databaseId 不能为空")
		ctx.JSON(http.StatusOK,gin.H{
			"code":500,
			"msg":"databaseId 不能为空",
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

	id, _ := strconv.Atoi(databaseId)
	db := connectionMaps[id]
	schemaInfo:=make(map[string]interface{})
	err := db.Raw("select * from information_schema.SCHEMATA where SCHEMA_NAME=?", schema).Scan(&schemaInfo)
	if err!=nil {
		log.Error("查询发生异常")
	}
	ctx.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"success",
		"data":schemaInfo,
	})
}





func Login(ctx *gin.Context) {
	userName := ctx.Query("userName")
	password := ctx.Query("password")
	result := make(map[string]string)
	if strings.EqualFold(userName, "admin") && strings.EqualFold(password, "123") {
		result["code"] = "200"
		result["msg"] = "成功"
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "登录成功",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 301,
		"msg":  "登录失败",
	})
}

//type DataSource struct {
//	Host string `json:"host"`
//	Port int `json:"port"`
//	Schema string `json:"schema"`
//	UserName string `json:"userName"`
//	Password string `json:"password"`
//}

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
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "链接失败",
			"data": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "链接成功",
	})
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
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "链接失败",
			"data": err2,
		})
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
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "保存数据库信息失败",
		})
		return
	}
	log.Info("保存数据库信息成功。")

	redisHelper := global.GetGlobal().RedisConnect
	json, _ := json.Marshal(myJdbc)

	result, err := redisHelper.Set(context.Background(), fmt.Sprintf("datasource_%d", dataConnect.Id), json, 0).Result()
	if err != nil {
		log.Error("缓存到数据失败", err)
		return
	}
	log.Info("缓存成功", result)

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "创建数据库链接成功",
	})

}

func GetRedisCache(ctx *gin.Context) {
	redis := global.GetGlobal().RedisConnect

	result, err := redis.Get(context.Background(), "datasource_1").Result()
	if err != nil {
		log.Error("获取缓存数据失败", err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": result,
	})
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
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": conns,
	})
}
