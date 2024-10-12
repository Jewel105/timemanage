package main

import (
	"fmt"
	"gin_study/config"
	"gin_study/gen/models"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open(config.MysqlDsn), &gorm.Config{})
	if err != nil {
		fmt.Println("mysql connect error", err.Error())
	}
	if db.Error != nil {
		fmt.Println("database error", db.Error)
	}
	e := db.AutoMigrate(&models.User{}, &models.Task{}, &models.Category{})
	if e != nil {
		fmt.Println("mysql migrate error", e.Error())
	}
	g := gen.NewGenerator(gen.Config{
		// 会自动创建目录
		OutPath: "./gen/query",
		// gen.WithoutContext：启用WithContext模式
		// gen.WithDefaultQuery：生成一个全局Query对象Q
		// gen.WithQueryInterface：生成Query接口
		Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.UseDB(db)
	g.ApplyBasic(&models.User{}, &models.Task{}, &models.Category{})
	g.Execute()
}