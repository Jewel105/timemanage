package router

import (
	"gin_study/api"
	categoryapi "gin_study/api/category_api"
	taskapi "gin_study/api/task_api"
	userapi "gin_study/api/user_api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	r := gin.Default()
	r.Use(gin.LoggerWithConfig(api.LoggerToFile()))
	r.Use(api.Recover)
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world")
	})
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

	r.Run()
}
