package userapi

import (
	"gin_study/api"
	"gin_study/api/consts"
	"gin_study/gen/request"
	"gin_study/language"
	userservice "gin_study/service/user_service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Id Login
// @Summary 登录
// @Description 登录
// @Tags 用户API
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

// @Id Logout
// @Summary 退出登录
// @Description 退出登录
// @Tags 用户API
// @Accept  json
// @Produce application/json
// @Param token header string false "enjmcvhdwernxhcuvyudfdjfhkjxkjaoerpq"
// @success 200 boolean ture "success"
// @Router /common/user/session/logout [get]
func Logout(c *gin.Context) {
	userID := api.GetUserID(c)
	if userID == 0 {
		return
	}
	tokenID, ok := c.Get(consts.TOKEN_ID)

	lang := c.GetString(consts.LANG)

	if !ok {
		api.DealResponse(c, false, &consts.ApiErr{Code: consts.TOKEN_INVALID, Msg: language.GetLocale(lang, consts.TOKEN_INVALID)})
	}

	success, e := userservice.Logout(userID, tokenID.(uuid.UUID), lang)
	api.DealResponse(c, success, e)
}

// @Id Register
// @Summary 注册
// @Description 注册
// @Tags 用户API
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
// @Tags 用户API
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
// @Tags 用户API
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

// @Id GetInfo
// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags 用户API
// @Accept  json
// @Produce application/json
// @Param token header string false "enjmcvhdwernxhcuvyudfdjfhkjxkjaoerpq"
// @success 200 {object} response.UserInfo "success"
// @Router /common/user/session/info [get]
func GetInfo(c *gin.Context) {
	userID := api.GetUserID(c)
	if userID == 0 {
		return
	}
	info, e := userservice.GetInfo(userID)
	api.DealResponse(c, info, e)
}

// @Id EditUserInfo
// @Summary 编辑用户信息
// @Description 编辑用户信息
// @Tags 用户API
// @Accept  json
// @Produce application/json
// @Param token header string false "enjmcvhdwernxhcuvyudfdjfhkjxkjaoerpq"
// @success 200 boolean ture "success"
// @Router /common/user/session/edit [post]
func EditUserInfo(c *gin.Context) {
	userID := api.GetUserID(c)
	if userID == 0 {
		return
	}
	tokenID, ok := c.Get(consts.TOKEN_ID)

	lang := c.GetString(consts.LANG)

	if !ok {
		api.DealResponse(c, false, &consts.ApiErr{Code: consts.TOKEN_INVALID, Msg: language.GetLocale(lang, consts.TOKEN_INVALID)})
	}

	success, e := userservice.Logout(userID, tokenID.(uuid.UUID), lang)
	api.DealResponse(c, success, e)
}
