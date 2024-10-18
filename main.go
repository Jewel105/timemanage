package main

import (
	"flag"
	"gin_study/config"
	"gin_study/factory"
	"gin_study/gen/mysql"
	"gin_study/logger"
	"gin_study/router"
	"sync"
)

var env string

func init() {
	flag.StringVar(&env, "env", "dev", "Specify the environment: dev, pro")
}

// @title time manage
// @version 1.0
// @description time manage
// @termsOfService http://127.0.0.1

// @contact.name jewel

// @host 127.0.0.1:8081
// @BasePath /api/v1
func main() {
	flag.Parse()
	config.GetConfig(env)

	var wg sync.WaitGroup
	wg.Add(3) // 有3个并发任务需要等待

	// 启动 MySQL
	go func() {
		defer wg.Done() // 标记任务完成
		mysql.Start()
	}()

	// 启动 Redis
	go func() {
		defer wg.Done()
		factory.RedisStart()
	}()

	// 启动 Logger
	defer logger.Sync()
	go func() {
		defer wg.Done()
		logger.InitLogger(logger.LogConfig{
			FileName:   "./log/timemanage.log",
			MaxSize:    100,
			MaxAge:     30,
			MaxBackups: 100,
		})
	}()

	// 等待 MySQL ， Redis ，logger初始化完成
	wg.Wait()
	router.Start()
}
