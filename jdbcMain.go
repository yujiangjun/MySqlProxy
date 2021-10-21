package main

import (
	"MySqlProxy/jdbc"
	log "github.com/sirupsen/logrus"
	"go/types"
)

func main() {

	var metaService jdbc.MetaService
	metaService = new(jdbc.Meta)
	meta := metaService.GetMeta()
	log.Infof("用户名:%s,密码:%s,url:%s",meta.Username,meta.Password,meta.Url)
	db := jdbc.GetConnection(meta)
	log.Info(db)
	exec, err := db.Exec("use raytine")
	if err!=nil {
		log.Error("发生错误.",err)
	}
	log.Info(exec)
	//values:=[3]string{"11","zhangsan","张三"}
	//jdbc.InsertUser("insert into im_user(user_id,user_name,nick_name) values(?,?,?)",values,db)
	var users []jdbc.ImUser
	jdbc.SelectUsers("select user_id,user_name ,nick_name  from im_user",db,&users)
	log.Info("select success",users)

	var resultMap []map[string]types.Object
	err = db.Select(&resultMap, "select user_id,user_name ,nick_name  from im_user")
	if err!=nil {
		log.Info("查询失败",err)
	}
	for key, value := range resultMap {
		log.Info("key:%s,value:%s",key,value)
	}
}
