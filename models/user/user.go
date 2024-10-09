package userModel

import (
	"gin_study/dao"
	categoryModel "gin_study/models/category"
	taskModel "gin_study/models/task"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string                   `gorm:"column:name;size:64" json:"name" binding:"required"`
	Password   string                   `gorm:"column:password;size:64" json:"password" binding:"required"`
	Tasks      []taskModel.Task         `json:"tasks"`
	Categories []categoryModel.Category `json:"categories"`
}

func GetInfoById(id int) (*User, error) {
	user := User{}
	result := dao.DB.First(&user, id)
	return &user, result.Error
}

func GetInfoByName(name string) (*User, error) {
	user := User{}
	result := dao.DB.Where("name = ?", name).First(&user)
	return &user, result.Error
}

func GetAll() ([]*User, error) {
	var users []*User
	result := dao.DB.Find(&users)
	return users, result.Error
}

// create or update
func SaveUser(user *User) error {
	if user.ID == 0 {
		result := dao.DB.Create(user)
		return result.Error
	}
	result := dao.DB.Model(user).Updates(user)
	return result.Error
}

func Delete(id int) error {
	result := dao.DB.Delete(&User{}, id)
	return result.Error
}
