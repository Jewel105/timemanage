package dao

import (
	"gin_study/config"
	"gin_study/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	DB, err = gorm.Open(mysql.Open(config.MysqlDsn), &gorm.Config{})
	if err != nil {
		logger.Error(map[string]interface{}{"mysql connect error": err.Error()})
	}
	if DB.Error != nil {
		logger.Error(map[string]interface{}{"database error": DB.Error})
	}

	// user := models.User{Name: "John Doe", Password: "123456"}
	// result := db.Create(&user)
	// if result.Error != nil {
	// 	// 处理创建错误
	// 	logger.Error(map[string]interface{}{"mysql Create table error": err.Error()})
	// }
}
