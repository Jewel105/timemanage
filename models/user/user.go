package userModel

import (
	"gin_study/dao"
	"time"
)

type User struct {
	Id          int64     `gorm:"column:id;primarykey;autoIncrement" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	Password    string    `gorm:"column:password" json:"password"`
	CreatedTime time.Time `gorm:"column:create_time;autoCreateTime:milli" json:"created_time"`
	UpdatedTime time.Time `gorm:"column:update_time;autoUpdateTime:milli" json:"updated_time"`
}

func GetInfoById(id int) (*User, error) {
	user := User{}
	result := dao.DB.First(&user, id)
	return &user, result.Error
}

func GetAll() ([]*User, error) {
	var users []*User
	result := dao.DB.Find(&users)
	return users, result.Error
}

// create or update
func SaveUser(user *User) error {
	if user.Id == 0 {
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
