package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Db *Database `mapstructure:"database"`
	RedisConfig *RedisConfig `mapstructure:"redis"`
}

type Database struct {
	SqlType string `mapstructure:"sqlType"`
	Url string `mapstructure:"url"`
	UserName string `mapstructure:"userName"`
	Password string `mapstructure:"password"`
}

type RedisConfig struct {
	Addr string `mapstructure:"url"`
	Password string `mapstructure:"password"`
}

var config *Config
func InitConfig() *Config{
	viper.SetConfigType("yaml")
	path, _ := os.Getwd()
	log.Info("路径:",path+"/config/jdbc.yml")
	viper.SetConfigFile(path+"/config/jdbc.yml")

	err := viper.ReadInConfig()
	if err!=nil {
		log.Error("加载数据库配置文件失败",err)
	}
	db:=&Database{}
	redis:=&RedisConfig{}
	 config :=&Config{
		 Db: db,
		 RedisConfig: redis,
	 }
	err = viper.Unmarshal(config)
	if err!=nil {
		log.Error("解析数据库配置文件失败",err)
	}
	log.Info("配置参数解析成功",&config)
	return config
}

func GetConfig() *Config {
	return config
}