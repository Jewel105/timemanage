package userapi

import (
	"gin_study/api"
	"gin_study/gen/request"
	userservice "gin_study/service/user_service"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	req := request.LoginRequest{}
	if !api.ParseJson(c, &req) {
		return
	}
	token, e := userservice.Login(&req)
	api.DealResponse(c, token, e)
}

func Register(c *gin.Context) {
	req := request.RegisterRequest{}
	if !api.ParseJson(c, &req) {
		return
	}
	userID, err := userservice.Register(&req)
	api.DealResponse(c, userID, err)
}
