package global

import (
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mySqlProxy/config"
	"time"
)

type Global struct {
	SqlConnect *gorm.DB
	RedisConnect *redis.Client

	Config *config.Config
}



func (g *Global) initSqlDb()  {

	if g.Config.Db.SqlType=="mysql"{
		connect, err := gorm.Open(mysql.Open(g.Config.Db.UserName+":"+g.Config.Db.Password+"@"+g.Config.Db.Url), &gorm.Config{})
		if err!=nil {
			log.Error("数据库链接异常",err)
		}
		db, _ := connect.DB()
		err = db.Ping()
		if err!=nil {
			log.Error("数据库链接异常",err)
		}
		g.SqlConnect=connect
	}
}

func (g *Global) initRedis()  {
	rdb := redis.NewClient(&redis.Options{
		Addr:         g.Config.RedisConfig.Addr,
		Password:     g.Config.RedisConfig.Password,
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
	g.RedisConnect=rdb
}

var global *Global

func NewGlobal() {
	 global =new(Global)
	 global.Config=config.InitConfig()
	 global.initSqlDb()
	 global.initRedis()
}

func GetGlobal() *Global {
	return global
}