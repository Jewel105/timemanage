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
	r.Use(api.RecordLog)
	r.Use(api.Recover)
	r.Use(api.SaveEquipmentID)

	// 引入swagger
	swaggerJson := getFileContent("./docs/swagger.json")
	knife.InitSwaggerKnife(r, swaggerJson)

	apiV1 := r.Group("/api/v1")
	common := apiV1.Group("/common")

	// 引入api
	userapi.AddRouter(common)
	systemapi.AddRouter(common)
	taskapi.AddRouter(apiV1)
	categoryapi.AddRouter(apiV1)

	r.Run(":" + config.Config.Server.Port)
}

func getFileContent(fpath string) string {
	bytes, err := os.ReadFile(fpath)
	if nil != err {

		logger.Error(zap.Any("Error file", err))

		return ""
	}

	return string(bytes)
}
