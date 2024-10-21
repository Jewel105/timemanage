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

	Name     string `gorm:"column:name;size:64" json:"name"`           // 用户名
	Password string `gorm:"column:password;size:128" json:"password"`  // 密码
	Email    string `gorm:"column:email;size:128;unique" json:"email"` // 邮箱
}

type Task struct {
	ID          int64          `gorm:"column:id;primarykey;autoIncrement" json:"id"`
	CreatedTime time.Time      `gorm:"column:create_time;autoCreateTime:milli" json:"createdTime"`
	UpdatedTime time.Time      `gorm:"column:update_time;autoUpdateTime:milli" json:"updatedTime"`
	DeleteTime  gorm.DeletedAt `gorm:"column:delete_time;index" json:"deleteTime"`

	UserID       int64  `gorm:"column:user_id" json:"userID"`                      // 创建该任务的用户ID
	Description  string `gorm:"column:description;size:200" json:"description"`    // 任务描述
	SpentTime    int64  `gorm:"column:spent_time" json:"spentTime"`                // 花费时间
	CategoryPath string `gorm:"column:category_path;size:128" json:"categoryPath"` // 任务路径上的所有ID
	StartTime    int64  `gorm:"column:start_time" json:"startTime"`                // 任务开始时间
	EndTime      int64  `gorm:"column:end_time" json:"endTime"`                    // 任务结束时间
	CategoryID   int64  `gorm:"column:category_id" json:"categoryID"`              // 任务所属分类ID
}

type Category struct {
	ID          int64          `gorm:"column:id;primarykey;autoIncrement" json:"id"`
	CreatedTime time.Time      `gorm:"column:create_time;autoCreateTime:milli" json:"createdTime"`
	UpdatedTime time.Time      `gorm:"column:update_time;autoUpdateTime:milli" json:"updatedTime"`
	DeleteTime  gorm.DeletedAt `gorm:"column:delete_time;index" json:"deleteTime"`

	Name     string `gorm:"column:name;size:64" json:"name"`   // 分类名称
	ParentID int64  `gorm:"column:parent_id" json:"parentID"`  // 上级分类ID
	UserID   int64  `gorm:"column:user_id" json:"userID"`      // 创建该分类的用户ID
	Path     string `gorm:"column:path;size:128" json:"path"`  // 分类路径
	Level    int    `gorm:"column:level;size:10" json:"level"` // 分类等级
}

type Equipment struct {
	ID          int64          `gorm:"column:id;primarykey;autoIncrement" json:"id"`
	CreatedTime time.Time      `gorm:"column:create_time;autoCreateTime:milli" json:"createdTime"`
	UpdatedTime time.Time      `gorm:"column:update_time;autoUpdateTime:milli" json:"updatedTime"`
	DeleteTime  gorm.DeletedAt `gorm:"column:delete_time;index" json:"deleteTime"`

	Vender           string `gorm:"column:vender;size:128" json:"vender"`                      // 供应商
	Type             string `gorm:"column:type;size:128" json:"type"`                          // 设备类型
	Sn               string `gorm:"column:sn;size:128" json:"sn"`                              // 序列号
	Imei1            string `gorm:"column:imei1;size:128" json:"imei1"`                        // IMEI1
	Imei0            string `gorm:"column:imei0;size:128" json:"imei0"`                        // IMEI0
	Os               string `gorm:"column:os;size:128" json:"os"`                              // 所属操作系统
	UserIDs          string `gorm:"column:user_ids;size:256" json:"userIDs"`                   // 所属所有用户ID
	IsPhysicalDevice int    `gorm:"column:is_physical_device;size:10" json:"isPhysicalDevice"` // 是否为物理设备
}

type FontLogs struct {
	ID          int64          `gorm:"column:id;primarykey;autoIncrement" json:"id"`
	CreatedTime time.Time      `gorm:"column:create_time;autoCreateTime:milli" json:"createdTime"`
	UpdatedTime time.Time      `gorm:"column:update_time;autoUpdateTime:milli" json:"updatedTime"`
	DeleteTime  gorm.DeletedAt `gorm:"column:delete_time;index" json:"deleteTime"`

	EquipmentID int64  `gorm:"column:equipment_id" json:"equipment_id"` // 所属设备ID
	UserID      int64  `gorm:"column:user_id" json:"userID"`            // 所属用户ID
	Version     string `gorm:"column:version;size:32" json:"version"`   // 版本
	Stack       string `gorm:"column:stack" json:"stack"`               // 堆栈
	Error       string `gorm:"column:error;size:256" json:"error"`      // 错误信息
}
