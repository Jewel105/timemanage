package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string     `gorm:"column:name;size:64" json:"name"`
	Password   string     `gorm:"column:password;size:64" json:"password"`
	Tasks      []Task     `json:"tasks"`
	Categories []Category `json:"categories"`
}

type Task struct {
	gorm.Model
	UserID      int64   `gorm:"column:user_id" json:"user_id"`                  // 创建该任务的用户ID
	Description string  `gorm:"column:description;size:200" json:"description"` // 任务描述
	SpentTime   float64 `gorm:"column:spent_time" json:"spent_time"`            // 花费时间
	CategoryID  int64   `gorm:"column:category_id" json:"category_id"`          // 任务所属分类ID
}

type Category struct {
	gorm.Model
	Name     string `gorm:"column:name;size:64" json:"name"`
	ParentID int64  `gorm:"column:parent_id" json:"parent_id"` // 上级分类ID
	UserID   int64  `gorm:"column:user_id" json:"user_id"`     // 创建该分类的用户ID
	Path     string `gorm:"column:path;size:64" json:"path"`   // 分类路径
}
