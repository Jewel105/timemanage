package systemapi

import (
	"github.com/gin-gonic/gin"
)

func AddRouter(router *gin.RouterGroup) {
	system := router.Group("/system")
	{
		system.POST("/register/equipment", RegisterEquipment) // 记录前端错误日志
		system.POST("/log/error", LogError)                   // 记录前端错误日志
	}
}
