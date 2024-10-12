package api

import (
	"gin_study/gen/request"
	"gin_study/service"

	"github.com/gin-gonic/gin"
)

type UserApi struct{}

func (u UserApi) Login(c *gin.Context) {
	req := request.LoginRequest{}
	if !ParseJson(c, &req) {
		return
	}
	token, e := service.UserService{}.Login(&req)
	DealResponse(c, token, e, LOGIN_FAILED)
}

func (u UserApi) Register(c *gin.Context) {
	req := request.RegisterRequest{}
	if !ParseJson(c, &req) {
		return
	}
	userID, err := service.UserService{}.Register(&req)
	DealResponse(c, userID, err, REGISTER_FAILED)
}
