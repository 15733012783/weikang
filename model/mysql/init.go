package mysql

import (
	"fmt"
	"github.com/15733012783/weikang/nacos"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InItMysql(c func(db *gorm.DB) (interface{}, error)) (interface{}, error) {
	nac := nacos.RpcNac.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		nac.Root, nac.Password, nac.Host, nac.Port, nac.Database)
	open, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	db, err := open.DB()
	if err != nil {
		return nil, err
	}
	s, err := c(open)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return s, err
}

func InitTable() {
	_, err := InItMysql(func(db *gorm.DB) (interface{}, error) {
		err := db.AutoMigrate()
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	if err != nil {
		return
	}
}
