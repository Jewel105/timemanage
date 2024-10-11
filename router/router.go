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

	order := r.Group("/order")
	order.GET("/list", api.OrderController{}.GetList)
	order.GET("/info", api.OrderController{}.GetInfo)

	r.Run()
}
