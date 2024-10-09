package api

import (
	"gin_study/factory"
	userModel "gin_study/models/user"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) Login(c *gin.Context) {
	user := userModel.User{}
	if !ParseJson(c, &user) {
		return
	}

	dbUser, e := userModel.GetInfoByName(user.Name)
	if e != nil {
		ReturnResponse(c, LOGIN_FAILED, "User and password are incorrect.")
		return
	}
	userPass := factory.Md5Hash(user.Password)
	if userPass != dbUser.Password {
		ReturnResponse(c, LOGIN_FAILED, "User and password are incorrect.")
		return
	}

	// get token
	token, e := factory.CreateToken(user.Name, 1)
	DealResponse(c, token, e)
}

func (u UserController) Register(c *gin.Context) {
	user := userModel.User{}
	if !ParseJson(c, &user) {
		return
	}
	dbUser, _ := userModel.GetInfoByName(user.Name)
	if dbUser != nil {
		ReturnResponse(c, CLIENT_ERROR, "User already exists.")
		return
	}
	user.Password = factory.Md5Hash(user.Password)
	err := userModel.SaveUser(&user)
	DealResponse(c, user.ID, err)
}
