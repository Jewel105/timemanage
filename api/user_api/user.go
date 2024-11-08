package userapi

import (
	"gin_study/api"
	"gin_study/api/consts"
	"gin_study/gen/request"
	userservice "gin_study/service/user_service"

	"github.com/gin-gonic/gin"
)

// @Id Login
// @Summary 登录
// @Description 登录
// @Tags COMMON API
// @Accept  json
// @Produce application/json
// @Param equipment header string false "3425243"
// @Param req body request.LoginRequest true "Json"
// @Success 200 {string} string "token"
// @Router  /common/user/login [post]
func Login(c *gin.Context) {
	req := request.LoginRequest{}
	if !api.ParseJson(c, &req) {
		return
	}
	equipmentID := c.GetInt64(consts.EQUIPMENT_ID)
	lang := c.GetString(consts.LANG)

	token, e := userservice.Login(equipmentID, &req, lang)
	api.DealResponse(c, token, e)
}

// @Id Register
// @Summary 注册
// @Description 注册
// @Tags COMMON API
// @Accept  json
// @Produce application/json
// @Param req body request.RegisterRequest true "Json"
// @Success 200 {int} int "1"
// @Router  /common/user/register [post]
func Register(c *gin.Context) {
	req := request.RegisterRequest{}
	if !api.ParseJson(c, &req) {
		return
	}
	lang := c.GetString(consts.LANG)

	userID, err := userservice.Register(&req, lang)
	api.DealResponse(c, userID, err)
}

// @Id SendCode
// @Summary 发送验证码
// @Description 发送验证码
// @Tags COMMON API
// @Accept json
// @Produce application/json
// @Param req body request.SendCodeRequest true "Json"
// @success 200 boolean ture "success"
// @Router  /common/user/send/code [post]
func SendCode(c *gin.Context) {
	req := request.SendCodeRequest{}
	if !api.ParseJson(c, &req) {
		return
	}
	lang := c.GetString(consts.LANG)

	err := userservice.SendCode(&req, lang)
	api.DealResponse(c, true, err)
}

// @Id ForgetPassword
// @Summary 忘记密码
// @Description 忘记密码
// @Tags COMMON API
// @Accept json
// @Produce application/json
// @Param req body request.RegisterRequest true "Json"
// @Success 200 {int} int "1"
// @Router  /common/user/forget/password [post]
func ForgetPassword(c *gin.Context) {
	req := request.RegisterRequest{}
	if !api.ParseJson(c, &req) {
		return
	}
	lang := c.GetString(consts.LANG)

	userID, err := userservice.ForgetPassword(&req, lang)
	api.DealResponse(c, userID, err)
}
