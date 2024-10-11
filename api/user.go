package api

import (
	"gin_study/factory"
	"gin_study/gen/models"
	"gin_study/gen/query"
	"gin_study/gen/request"

	"github.com/gin-gonic/gin"
)

type UserApi struct{}

func (u UserApi) Login(c *gin.Context) {
	req := request.LoginRequest{}
	if !ParseJson(c, &req) {
		return
	}
	user, err := query.User.Where(query.User.Name.Eq(req.Name)).First()
	if err != nil {
		ReturnResponse(c, LOGIN_FAILED, "User and password are incorrect.")
		return
	}
	reqPass := factory.Md5Hash(req.Password)
	if reqPass != user.Password {
		ReturnResponse(c, LOGIN_FAILED, "User and password are incorrect.")
		return
	}
	// get token
	token, e := factory.CreateToken(user.Name, user.ID)
	DealResponse(c, token, e)
}

func (u UserApi) Register(c *gin.Context) {
	req := request.RegisterRequest{}
	if !ParseJson(c, &req) {
		return
	}
	userExists, _ := query.User.Where(query.User.Name.Eq(req.Name)).First()
	if userExists != nil {
		ReturnResponse(c, CLIENT_ERROR, "User already exists.")
		return
	}
	user := models.User{
		Name:     req.Name,
		Password: factory.Md5Hash(req.Password),
	}
	tx := query.Q.Begin()
	err := query.User.Save(&user)
	if err != nil {
		err = tx.Rollback()
	} else {
		err = tx.Commit()
	}
	DealResponse(c, user.ID, err)
}
