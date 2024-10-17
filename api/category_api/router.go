package categoryapi

import (
	"gin_study/api"

	"github.com/gin-gonic/gin"
)

func AddRouter(router *gin.RouterGroup) {
	categories := router.Group("/categories")
	categories.Use(api.VerifyToken)
	{
		categories.GET("/list", GetList)
		categories.POST("/save", SaveCategory)
		categories.POST("/delete/:id", DeleteCategory)
	}
}
