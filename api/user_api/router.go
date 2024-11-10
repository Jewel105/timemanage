package userapi

import (
	"gin_study/api"

	"github.com/gin-gonic/gin"
)

func AddRouter(router *gin.RouterGroup) {
	user := router.Group("/user")
	{
		user.POST("/login", Login)
		user.POST("/register", Register)
		user.POST("/send/code", SendCode)
		user.POST("/forget/password", ForgetPassword)
	}

	log := user.Group("/log")
	log.Use(api.VerifyToken)
	log.GET("/out", Logout)

}
