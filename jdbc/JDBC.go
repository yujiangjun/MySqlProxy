package jdbc

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
import (
	log "github.com/sirupsen/logrus"
)

type MyJdbc struct {
	connect sql.DB
	stmt    sql.Stmt
	rows    sql.Rows
}

func getConnection(meta Meta) *sql.DB {
	connect, err := sql.Open("mysql", meta.username+":"+meta.password+"@"+meta.url)
	if err != nil {
		log.Error("链接失败")
	}
	defer connect.Close()
	return connect
}
