package router

import (
	"gin_study/api"
	categoryapi "gin_study/api/category_api"
	systemapi "gin_study/api/system_api"
	taskapi "gin_study/api/task_api"
	userapi "gin_study/api/user_api"
	"gin_study/config"
	"gin_study/logger"
	"os"

	knife "gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// swag init --parseDependency --parseInternal ---swagger json生成命令
// export PATH=$(go env GOPATH)/bin:$PATH // 找不到swag
func Start() {
	r := gin.Default()
	// 中间件
	r.Use(api.RecordLog)
	r.Use(api.Recover)
	r.Use(api.SaveEquipmentID)
	// 查看http协议
	// r.Use(func(c *gin.Context) {
	// 	protocol := c.Request.Proto
	// 	fmt.Printf("Protocol: %s\n", protocol)
	// 	c.Next()
	// })

	// 引入swagger
	swaggerJson := getFileContent("./docs/swagger.json")
	knife.InitSwaggerKnife(r, swaggerJson)

	// 引入api
	apiV1 := r.Group("/api/v1")
	common := apiV1.Group("/common")
	userapi.AddRouter(common)
	systemapi.AddRouter(common)
	taskapi.AddRouter(apiV1)
	categoryapi.AddRouter(apiV1)

	// 是否启用 H2C（HTTP/2 Cleartext）
	r.UseH2C = config.Config.Server.EnableH2C
	// 是否开启 HTTPS
	if config.Config.Server.EnableSSL {
		err := r.RunTLS(":"+config.Config.Server.Port, config.Config.Server.Certificate.Cert, config.Config.Server.Certificate.Key)
		if err != nil {
			panic(err)
		}
	} else {
		err := r.Run(":" + config.Config.Server.Port)
		if err != nil {
			panic(err)
		}
	}
}

func getFileContent(fpath string) string {
	bytes, err := os.ReadFile(fpath)
	if nil != err {

		logger.Error(zap.Any("Error file", err))

		return ""
	}

	return string(bytes)
}
