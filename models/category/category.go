package categoryModel

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name     string `gorm:"column:name;size:64" json:"name"`
	ParentID int64  `gorm:"column:parent_id" json:"parent_id"` // 上级分类ID
	UserID   int64  `gorm:"column:user_id" json:"user_id"`     // 创建该分类的用户ID
	Path     string `gorm:"column:path;size:64" json:"path"`   // 分类路径
}
