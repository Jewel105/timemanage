package api

import (
	"fmt"
	userModel "gin_study/models/user"
	"gin_study/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) GetInfo(c *gin.Context) { //-
	json := make(map[string]interface{}) //-
	err := c.BindJSON(&json)             //-
	if err != nil {                      //-
		ReturnResponse(c, SYSTEM_ERROR, err.Error()) //-
		return                                       //-
	} //-
	token := json["token"]
	if token == nil {
		ReturnResponse(c, SYSTEM_ERROR, "token不能为空")
		return
	}
	uu, ett := util.DecodeToken(token.(string)) //+
	if ett != nil {
		ReturnResponse(c, TOKEN_INVALID, ett.Error())
		return
	}
	ReturnResponse(c, SUCCESS, uu) //-
	fmt.Println(uu.UserName)       //+
}

func (u UserController) GetList(c *gin.Context) {
	users, _ := userModel.GetAll()
	ReturnResponse(c, SUCCESS, users)

}

func (u UserController) Login(c *gin.Context) {
	user := userModel.User{}
	err := c.BindJSON(&user)
	if err != nil {
		ReturnResponse(c, SYSTEM_ERROR, err.Error())
		return
	}
	// get token
	token, e := util.CreateToken(user.Name, 1)
	if e != nil {
		ReturnResponse(c, SYSTEM_ERROR, e.Error())
		return
	}
	// userModel.SaveUser(&user)
	ReturnResponse(c, SUCCESS, token)

}

func (u UserController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	userModel.Delete(id)
	ReturnResponse(c, SYSTEM_ERROR, id)
}
