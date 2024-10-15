package router

import (
	"fmt"
	"gin_study/api"
	categoryapi "gin_study/api/category_api"
	taskapi "gin_study/api/task_api"
	userapi "gin_study/api/user_api"
	"gin_study/config"
	"os"

	knife "gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife"

	"github.com/gin-gonic/gin"
)

// swag init --parseDependency --parseInternal
func Start() {
	r := gin.Default()
	r.Use(gin.LoggerWithConfig(api.LoggerToFile()))
	r.Use(api.Recover)

	swaggerJson := getFileContent("./docs/swagger.json")
	knife.InitSwaggerKnife(r, swaggerJson)

	apiV1 := r.Group("/api/v1")
	common := apiV1.Group("/common")
	user := common.Group("/user")
	{
		user.POST("/login", userapi.Login)
		user.POST("/register", userapi.Register)
	}

	tasks := apiV1.Group("/tasks")
	tasks.Use(api.VerifyToken)
	{
		tasks.GET("/list", taskapi.GetList)
		tasks.POST("/save", taskapi.SaveTask)
		tasks.POST("/delete/:id", taskapi.DeleteTask)
	}

	categories := apiV1.Group("/categories")
	categories.Use(api.VerifyToken)
	{
		categories.GET("/list", categoryapi.GetList)
		categories.POST("/save", categoryapi.SaveCategory)
		categories.POST("/delete/:id", categoryapi.DeleteCategory)
	}

	r.Run(":" + config.Config.Server.Port)
}

func getFileContent(fpath string) string {
	bytes, err := os.ReadFile(fpath)
	if nil != err {
		fmt.Errorf(" %s getFileBase64 error: %v", fpath, err)
		return ""
	}

	return string(bytes)
}
