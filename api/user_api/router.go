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

	session := user.Group("/session")
	session.Use(api.VerifyToken)
	{
		session.GET("/logout", Logout)
		session.GET("/info", GetInfo)
		session.POST("/edit", EditUserInfo)
	}
}
