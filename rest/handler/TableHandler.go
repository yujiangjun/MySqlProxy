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

func GetContext( context *gin.Context){
	databaseId, ok := context.GetQuery("databaseId")
	if !ok  {
		log.Info("缺少tableName")
		return
	}
	log.Info("databaseId:%s",databaseId)
	id, _ := strconv.Atoi(databaseId)
	log.Info(connectionMaps[id])
	var tables []map[string]interface{}
	err := connectionMaps[id].Raw( "SELECT COLUMN_NAME fName,column_comment fDesc,DATA_TYPE dataType, IS_NULLABLE isNull,IFNULL(CHARACTER_MAXIMUM_LENGTH,0) sLength FROM information_schema.columns ").Scan(&tables)
	if err!=nil {
		log.Error("查询失败",err)
	}
	context.JSON(http.StatusOK,tables)
}



func GetTables(ctx *gin.Context) {
	databaseId, ok := ctx.GetQuery("databaseId")
	if !ok {
		log.Error("为获取到参数schema")
		return
	}
	id, _ := strconv.Atoi(databaseId)

	//var tables *[]jdbc.Tables
	//maps :=[...]map[string]interface{}
	conn:= connectionMaps[id]
	log.Info("链接信息",conn)
	var tables []map[string]interface{}
	err := connectionMaps[id].Raw("select TABLE_CATALOG,TABLE_SCHEMA,TABLE_NAME,TABLE_TYPE,ENGINE,VERSION,ROW_FORMAT,TABLE_ROWS,DATA_LENGTH,CREATE_TIME,UPDATE_TIME,TABLE_COLLATION,TABLE_COMMENT from information_schema.TABLES").Scan(&tables)
	//err := connectionMaps[id].Raw("select TABLE_CATALOG,TABLE_SCHEMA,TABLE_NAME from information_schema.TABLES").Scan(&tables)
	if err!=nil {
		log.Error("查询发生异常.",err)
	}
	ctx.JSON(http.StatusOK,tables)
}

func Login(ctx *gin.Context) {
	userName := ctx.Query("userName")
	password := ctx.Query("password")
	result:= make(map[string]string)
	if  strings.EqualFold(userName,"admin") && strings.EqualFold(password,"123"){
		result["code"]="200"
		result["msg"]="成功"
		ctx.JSON(http.StatusOK,gin.H{
			"code":200,
			"msg":"登录成功",
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"code":301,
		"msg":"登录失败",
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

	dataSource :=dto.DataSource{}
	ctx.BindJSON(&dataSource)


	log.Info("参数:",dataSource)
	log.Info("拼接的url:",fmt.Sprintf("tcp(%s:%d)/%s?charset=utf8&parseTime=true",dataSource.Host,dataSource.Port,dataSource.Schema))
	meta := jdbc.MyJdbc{
		Url:      fmt.Sprintf("tcp(%s:%d)/%s?charset=utf8&parseTime=true",dataSource.Host,dataSource.Port,dataSource.Schema),
		Username: dataSource.UserName,
		Password: dataSource.Password,
	}
	connection:=meta.GetConnection()
	db,_ := connection.DB()
	err := db.Ping()
	if err !=nil {
		ctx.JSON(http.StatusOK,gin.H{
			"code":500,
			"msg":"链接失败",
			"data":err,
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"链接成功",
	})
}

func CreateConnect(ctx *gin.Context) {
	dataSource := dto.DataSource{}
	ctx.BindJSON(&dataSource)
	log.Info("dataSource:",dataSource)
	myJdbc := jdbc.MyJdbc{
		Username: dataSource.UserName,
		Password: dataSource.Password,
		Url:      fmt.Sprintf("tcp(%s:%d)/%s?charset=utf8&parseTime=true", dataSource.Host, dataSource.Port, dataSource.Schema),
	}
	connection := myJdbc.GetConnection()
	db, _ := connection.DB()
	err2 := db.Ping()
	if err2 !=nil{
		ctx.JSON(http.StatusOK,gin.H{
			"code":500,
			"msg":"链接失败",
			"data":err2,
		})
		return
	}
	metaService := new(jdbc.Meta)

	log.Info("记录链接信息",metaService.GetMeta())
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
	connectionMaps[dataConnect.Id]=connection

	if create.Error!=nil {
		log.Error("保存数据库信息失败：",create.Error)
		ctx.JSON(http.StatusOK,gin.H{
			"code":500,
			"msg":"保存数据库信息失败",
		})
		return
	}
	log.Info("保存数据库信息成功。")

	redisHelper := global.GetGlobal().RedisConnect
	json, _ := json.Marshal(myJdbc)

	result, err := redisHelper.Set(context.Background(), fmt.Sprintf("datasource_%d",dataConnect.Id), json, 0).Result()
	if err!=nil {
		log.Error("缓存到数据失败",err)
		return
	}
	log.Info("缓存成功",result)

	ctx.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"创建数据库链接成功",
	})

}

func GetRedisCache(ctx *gin.Context)  {
	redis := global.GetGlobal().RedisConnect

	result, err := redis.Get(context.Background(), "datasource*").Result()
	if err!=nil {
		log.Error("获取缓存数据失败",err)
	}
	ctx.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"ok",
		"data":result,
	})
}

func InitLoadingConnection2Redis() {
	db:= global.GetGlobal().SqlConnect
	redisDb := global.GetGlobal().RedisConnect
	var dataConnect []*entity.DataConnect
	db.Table("data_connect").Find(&dataConnect)
	for _,value := range dataConnect{

		connectionMaps[value.Id]=jdbc.MyJdbc{
			Username: value.UserName,
			Password: value.Password,
			Url:      value.Url,
		}.GetConnection()
		json, _ := json.Marshal(value)
		err := redisDb.Set(context.Background(), fmt.Sprintf("datasource_%d", value.Id), json, 0).Err()
		if err!=nil {
			log.Error("存入缓存缓存错误",err)
		}
	}
	log.Info("加载链接到缓存成功")
}