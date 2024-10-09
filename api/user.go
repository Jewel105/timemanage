package api

import (
	"gin_study/factory"
	userModel "gin_study/models/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) GetInfo(c *gin.Context) {
	json := make(map[string]interface{})
	err := c.BindJSON(&json)
	if err != nil {
		ReturnResponse(c, SYSTEM_ERROR, err.Error())
		return
	}
	token := json["token"]
	if token == nil {
		ReturnResponse(c, SYSTEM_ERROR, "token不能为空")
		return
	}
	uu, ett := factory.DecodeToken(token.(string))
	if ett != nil {
		ReturnResponse(c, TOKEN_INVALID, ett.Error())
		return
	}
	ReturnResponse(c, SUCCESS, uu)
}

func (u UserController) GetList(c *gin.Context) {
	users, _ := userModel.GetAll()
	ReturnResponse(c, SUCCESS, users)
}

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

func (u UserController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	userModel.Delete(id)
	ReturnResponse(c, SYSTEM_ERROR, id)
}
