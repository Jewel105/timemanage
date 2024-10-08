package taskModel

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	UserID      int64   `gorm:"column:user_id" json:"user_id"`                  // 创建该任务的用户ID
	Description string  `gorm:"column:description;size:200" json:"description"` // 任务描述
	SpentTime   float64 `gorm:"column:spent_time" json:"spent_time"`            // 花费时间
	CategoryID  int64   `gorm:"column:category_id" json:"category_id"`          // 任务所属分类ID
}
