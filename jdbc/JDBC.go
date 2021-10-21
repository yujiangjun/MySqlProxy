package jdbc

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)
import (
	log "github.com/sirupsen/logrus"
)

type MyJdbc struct {
	connect sql.DB
	stmt    sql.Stmt
	rows    sql.Rows
}

func GetConnection(meta Meta) *sqlx.DB {
	connect, err := sqlx.Open("mysql", meta.Username+":"+meta.Password+"@"+meta.Url)
	if err != nil {
		log.Error("链接失败")
	}
	return connect
}

func InsertUser(sql string,values [3]string,db *sqlx.DB)  {
	exec, err := db.Exec(sql, values[0], values[1], values[2])
	if err !=nil {
		log.Error("插入失败",err)
		return
	}
	id, err := exec.LastInsertId()
	if err!=nil{
		log.Error("exec failed",err)
	}
	log.Info("insert succ",id)
}

func SelectUsers(sql string,db *sqlx.DB,users *[]ImUser)  {
	err := db.Select(users, sql)
	if err !=nil {
		log.Errorf("查询失败,%s",err)
	}
}