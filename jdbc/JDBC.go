package jdbc

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connect interface {
	GetConnection() *gorm.DB
}
type MyJdbc struct {
	Username string
	Password string
	Url      string
}

func (jdbc MyJdbc) GetConnection() *gorm.DB {
	connect, err := gorm.Open(mysql.Open(jdbc.Username+":"+jdbc.Password+"@"+jdbc.Url), &gorm.Config{})
	if err != nil {
		log.Error("链接失败")
	}
	return connect
}