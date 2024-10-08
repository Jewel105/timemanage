package controllers

import (
	userModel "gin_study/models/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) GetInfo(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	user, _ := userModel.GetInfoById(id)
	ReturnSuccess(c, "0", "success", user, 3)
}

func (u UserController) GetList(c *gin.Context) {
	users, _ := userModel.GetAll()
	ReturnSuccess(c, "0", "success", users, 3)
}

func (u UserController) SaveUser(c *gin.Context) {
	user := userModel.User{}
	err := c.BindJSON(&user)
	if err != nil {
		ReturnError(c, "4001", gin.H{"err": err.Error()})
		return
	}
	userModel.SaveUser(&user)
	ReturnSuccess(c, "0", "success", user.ID, 3)
}

func (u UserController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	userModel.Delete(id)
	ReturnSuccess(c, "0", "success", id, 3)
}
