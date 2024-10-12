package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          int64          `gorm:"column:id;primarykey;autoIncrement" json:"id"`
	CreatedTime time.Time      `gorm:"column:create_time;autoCreateTime:milli" json:"createdTime"`
	UpdatedTime time.Time      `gorm:"column:update_time;autoUpdateTime:milli" json:"updatedTime"`
	DeleteTime  gorm.DeletedAt `gorm:"column:delete_time;index" json:"deleteTime"`

	Name       string     `gorm:"column:name;size:64" json:"name"`
	Password   string     `gorm:"column:password;size:128" json:"password"`
	Tasks      []Task     `json:"tasks"`
	Categories []Category `json:"categories"`
}

type Task struct {
	ID          int64          `gorm:"column:id;primarykey;autoIncrement" json:"id"`
	CreatedTime time.Time      `gorm:"column:create_time;autoCreateTime:milli" json:"createdTime"`
	UpdatedTime time.Time      `gorm:"column:update_time;autoUpdateTime:milli" json:"updatedTime"`
	DeleteTime  gorm.DeletedAt `gorm:"column:delete_time;index" json:"deleteTime"`

	UserID      int64  `gorm:"column:user_id" json:"userID"`                   // 创建该任务的用户ID
	Description string `gorm:"column:description;size:200" json:"description"` // 任务描述
	SpentTime   int64  `gorm:"column:spent_time" json:"spentTime"`             // 花费时间
	CategoryID  int64  `gorm:"column:category_id" json:"categoryID"`           // 任务所属分类ID
	StartTime   int64  `gorm:"column:start_time" json:"startTime"`
	EndTime     int64  `gorm:"column:end_time" json:"endTime"`
}

type Category struct {
	ID          int64          `gorm:"column:id;primarykey;autoIncrement" json:"id"`
	CreatedTime time.Time      `gorm:"column:create_time;autoCreateTime:milli" json:"createdTime"`
	UpdatedTime time.Time      `gorm:"column:update_time;autoUpdateTime:milli" json:"updatedTime"`
	DeleteTime  gorm.DeletedAt `gorm:"column:delete_time;index" json:"deleteTime"`

	Name     string `gorm:"column:name;size:64" json:"name"`
	ParentID int64  `gorm:"column:parent_id" json:"parentID"`  // 上级分类ID
	UserID   int64  `gorm:"column:user_id" json:"userID"`      // 创建该分类的用户ID
	Path     string `gorm:"column:path" json:"path"`           // 分类路径
	Level    int    `gorm:"column:level;size:10" json:"level"` // 分类等级
}