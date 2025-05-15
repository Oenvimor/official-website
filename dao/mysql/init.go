package mysql

import (
	"cqupt_hub/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Init(cfg *setting.Mysql) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	//db.Use(dbresolver.Register(dbresolver.Config{
	//	// 主库配置
	//	Sources: []gorm.Dialector{mysql.Open(dsn)},
	//	// 从库配置
	//	Replicas: []gorm.Dialector{mysql.Open(cfg.SlaveDns)},
	//}))
	Migration(db)
	return db
}
