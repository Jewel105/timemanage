package models

import (
	"gin_study/dao"
	"gin_study/logger"
	categoryModel "gin_study/models/category"
	taskModel "gin_study/models/task"
	userModel "gin_study/models/user"
)

func init() {
	err := dao.DB.AutoMigrate(&userModel.User{}, &taskModel.Task{}, &categoryModel.Category{})
	if err != nil {
		logger.Error(map[string]interface{}{"mysql migrate error": err.Error()})
	}
}
