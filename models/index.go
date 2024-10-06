package models

import (
	"gin_study/dao"
	"gin_study/logger"
	userModel "gin_study/models/user"
)

func init() {
	err := dao.DB.AutoMigrate(&userModel.User{})
	if err != nil {
		logger.Error(map[string]interface{}{"mysql migrate error": err.Error()})
	}
}
