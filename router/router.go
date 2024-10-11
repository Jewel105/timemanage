package router

import (
	"gin_study/api"
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
		user.POST("/login", api.UserApi{}.Login)
		user.POST("/register", api.UserApi{}.Register)
	}

	tasks := apiV1.Group("/tasks")
	tasks.Use(api.VerifyToken)
	{
		tasks.GET("/list", api.TaskApi{}.GetList)
		tasks.POST("/save", api.TaskApi{}.SaveTask)
		tasks.POST("/delete/:id", api.TaskApi{}.DeleteTask)
	}

	categories := apiV1.Group("/categories")
	categories.Use(api.VerifyToken)
	{
		categories.GET("/list", api.CategoryApi{}.GetList)
		categories.POST("/save", api.CategoryApi{}.SaveCategory)
		categories.POST("/delete/:id", api.CategoryApi{}.DeleteCategory)
	}

	r.Run()
}
