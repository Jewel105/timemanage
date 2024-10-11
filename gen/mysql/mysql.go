package mysql

import (
	"gin_study/config"
	"gin_study/gen/query"
	"gin_study/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	db, err := gorm.Open(mysql.Open(config.MysqlDsn), &gorm.Config{})
	if err != nil {
		logger.Error(map[string]interface{}{"mysql connect error": err.Error()})
	}
	if db.Error != nil {
		logger.Error(map[string]interface{}{"database error": db.Error})
	}
	query.SetDefault(db)
}
