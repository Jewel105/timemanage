package main

import (
	"flag"
	"gin_study/config"
	"gin_study/factory"
	"gin_study/gen/mysql"
	"gin_study/router"
	"sync"
)

var env string

func init() {
	flag.StringVar(&env, "env", "dev", "Specify the environment: dev, pro")
}

// @title 这里写标题
// @version 1.0
// @description 这里写描述信息
// @termsOfService http://swagger.io/terms/

// @contact.name 这里写联系人信息
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 这里写接口服务的host
// @BasePath 这里写base path
func main() {
	flag.Parse()
	config.GetConfig(env)
	var wg sync.WaitGroup
	wg.Add(2) // 有两个并发任务需要等待

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

	// 等待 MySQL 和 Redis 初始化完成
	wg.Wait()
	router.Start()
}
