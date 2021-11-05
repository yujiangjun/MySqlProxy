package jdbc

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
import (
	log "github.com/sirupsen/logrus"
)

type MyJdbc struct {
	connect sql.DB
	stmt    sql.Stmt
	rows    sql.Rows
}

func GetConnection(meta Meta) *gorm.DB {
	connect, err := gorm.Open(mysql.Open(meta.Username+":"+meta.Password+"@"+meta.Url), &gorm.Config{})
	//connect, err := sqlx.Open("mysql", meta.Username+":"+meta.Password+"@"+meta.Url)
	if err != nil {
		log.Error("链接失败")
	}
	return connect
}