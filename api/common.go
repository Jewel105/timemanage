package api

import (
	"errors"
	"gin_study/api/consts"
	"strings"

	"github.com/gin-gonic/gin"
)

type JsonStruct struct {
	Code    string      `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

func ReturnResponse(c *gin.Context, code string, data interface{}) {
	json := &JsonStruct{
		Code:    code,
		Msg:     GetMsg(code),
		Data:    data,
		Success: code == "200",
	}
	c.JSON(200, json)
}

func DealResponse(c *gin.Context, data interface{}, err error) {
	if err != nil {
		var ApiErr *consts.ApiErr
		if errors.As(err, &ApiErr) {
			ReturnResponse(c, ApiErr.Code, err.Error())
		} else {
			ReturnResponse(c, consts.SYSTEM_ERROR, err.Error())
		}
		return
	}
	ReturnResponse(c, consts.SUCCESS, data)
}

// GetMsg 返回错误的消息解释
func GetMsg(code string) string {
	msg, ok := consts.MsgFlags[code]
	if ok {
		return strings.Replace(strings.ToLower(msg), "_", " ", -1)
	}
	return consts.MsgFlags[consts.SYSTEM_ERROR]
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
	userID, exists := c.Get(consts.USER_ID)
	if !exists {
		ReturnResponse(c, consts.TOKEN_INVALID, "user not found")
		return 0
	}
	return userID.(int64)
}
