package router

import (
	"fmt"
	"gin_study/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(gin.LoggerWithConfig(controllers.LoggerToFile()))
	r.Use(controllers.Recover)
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world")
	})
	user := r.Group("/user")
	user.GET("/list", controllers.UserController{}.GetList)
	user.GET("/info/:id", controllers.UserController{}.GetInfo)
	user.POST("/save", controllers.UserController{}.SaveUser)
	user.POST("/delete/:id", controllers.UserController{}.Delete)

	order := r.Group("/order")
	order.GET("/list", controllers.OrderController{}.GetList)
	order.GET("/info", controllers.OrderController{}.GetInfo)

	fmt.Println("fmt")
	return r
}
