package userapi

import (
	"github.com/gin-gonic/gin"
)

func AddRouter(router *gin.RouterGroup) {
	user := router.Group("/user")
	{
		user.POST("/login", Login)
		user.POST("/register", Register)
	}
}
