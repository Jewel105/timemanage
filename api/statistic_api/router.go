package statisticapi

import (
	"gin_study/api"

	"github.com/gin-gonic/gin"
)

func AddRouter(router *gin.RouterGroup) {
	categories := router.Group("/statistic")
	categories.Use(api.VerifyToken)
	{
		categories.POST("/pie", GetPieValue)
		categories.POST("/line", GetLineValue)
	}
}
