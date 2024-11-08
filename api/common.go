package api

import (
	"errors"
	"gin_study/api/consts"
	"gin_study/language"

	"github.com/gin-gonic/gin"
)

type JsonStruct struct {
	Code    string      `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

func ReturnResponse(c *gin.Context, code string, data interface{}) {
	lang := c.GetString(consts.LANG)
	json := &JsonStruct{
		Code:    code,
		Msg:     language.GetLocale(lang, code),
		Data:    data,
		Success: code == "200",
	}
	c.JSON(200, json)
}

func DealResponse(c *gin.Context, data interface{}, err error) {
	if err != nil {
		var apiErr *consts.ApiErr
		if errors.As(err, &apiErr) {
			ReturnResponse(c, apiErr.Code, apiErr.Error())
		} else {
			ReturnResponse(c, consts.SYSTEM_ERROR, err.Error())
		}
		return
	}
	ReturnResponse(c, consts.SUCCESS, data)
}

func ParseJson(c *gin.Context, obj any) bool {
	err := c.ShouldBindJSON(obj)
	if err != nil {
		ReturnResponse(c, consts.PARAMS_INVALID, err.Error())
		return false
	}
	return true
}

func ParseQuery(c *gin.Context, obj any) bool {
	err := c.ShouldBindQuery(obj)
	if err != nil {
		ReturnResponse(c, consts.PARAMS_INVALID, err.Error())
		return false
	}
	return true
}

func GetUserID(c *gin.Context) int64 {
	userID := c.GetInt64(consts.USER_ID)
	if userID == 0 {
		ReturnResponse(c, consts.TOKEN_INVALID, "user not found")
	}
	return userID
}
