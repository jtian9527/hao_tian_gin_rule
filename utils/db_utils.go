package utils

import (
	"fmt"
	"haotian_rule/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBCluster = GetConn()

func GetConn() *gorm.DB {
	conf, err := config.ParseConfig("./config/config.json")
	if err != nil {
		panic("读取配置文件失败，" + err.Error())
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&autocommit=1", conf.Database.User, conf.Database.Password, conf.Database.Host, conf.Database.Port, conf.Database.DbName)
	DBCluster, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("GetConn err, error=" + err.Error())
	}
	sqlDB, _ := DBCluster.DB()

	//设置连接池参数
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(20)
	return DBCluster
}
