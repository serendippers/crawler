package initialize

import (
	"crawler/global"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func BizMysql() {
	bizConfig := global.CONFIG.MySQL
	source := bizConfig.Username + ":" + bizConfig.Password + "@(" + bizConfig.Path + ")/" + bizConfig.DBName + "?" + bizConfig.Config
	if bizDB, err := gorm.Open("mysql", source); err != nil {
		global.LOG.Error("DEFAULTDB数据库启动异常", err)
	} else {
		global.BIZ_DB = bizDB
		global.BIZ_DB.DB().SetMaxIdleConns(global.CONFIG.MaxIdleConns)
		global.BIZ_DB.DB().SetMaxOpenConns(global.CONFIG.MaxOpenConns)
		global.BIZ_DB.LogMode(bizConfig.LogMode)
	}
}
