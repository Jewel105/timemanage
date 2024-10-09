package router

import (
	"fmt"
	"gin_study/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
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
		user.POST("/login", api.UserController{}.Login)
		user.POST("/register", api.UserController{}.Register)

		user.GET("/list", api.UserController{}.GetList)
		user.POST("/info", api.UserController{}.GetInfo)
		user.POST("/delete/:id", api.UserController{}.Delete)
	}

	order := r.Group("/order")
	order.GET("/list", api.OrderController{}.GetList)
	order.GET("/info", api.OrderController{}.GetInfo)

	fmt.Println("fmt")
	return r
}
