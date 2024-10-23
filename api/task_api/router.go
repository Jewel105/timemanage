package taskapi

import (
	"gin_study/api"

	"github.com/gin-gonic/gin"
)

func AddRouter(router *gin.RouterGroup) {
	tasks := router.Group("/tasks")
	tasks.Use(api.VerifyToken)
	{
		tasks.GET("/list", GetList)
		tasks.POST("/save", SaveTask)
		tasks.POST("/delete/:id", DeleteTask)
		tasks.GET("/last/time", GetLastEndTime)
	}
}
