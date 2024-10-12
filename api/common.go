package api

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type JsonStruct struct {
	Code    string      `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

const (
	SUCCESS       = "200"
	SYSTEM_ERROR  = "500"
	NETWORK_ERROR = "400"

	PARAMS_INVALID  = "300"
	TOKEN_INVALID   = "301"
	LOGIN_FAILED    = "302"
	REGISTER_FAILED = "303"
)

const (
	USER_ID = "userID"
)

var MsgFlags = map[string]string{
	SUCCESS:       "SUCCESS",
	SYSTEM_ERROR:  "SYSTEM_ERROR",
	NETWORK_ERROR: "NETWORK_ERROR",

	PARAMS_INVALID:  "PARAMS_INVALID",
	TOKEN_INVALID:   "TOKEN_INVALID",
	LOGIN_FAILED:    "LOGIN_FAILED",
	REGISTER_FAILED: "REGISTER_FAILED",
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

func DealResponse(c *gin.Context, data interface{}, err error, errCode ...string) {
	if err != nil {
		if len(errCode) > 0 {
			ReturnResponse(c, errCode[0], err.Error())
		} else {
			ReturnResponse(c, SYSTEM_ERROR, err.Error())
		}
		return
	}
	ReturnResponse(c, SUCCESS, data)
}

// GetMsg 返回错误的消息解释
func GetMsg(code string) string {
	msg, ok := MsgFlags[code]
	if ok {
		return strings.Replace(strings.ToLower(msg), "_", " ", -1)
	}
	return MsgFlags[SYSTEM_ERROR]
}

func ParseJson(c *gin.Context, obj any) bool {
	err := c.ShouldBindJSON(obj)
	if err != nil {
		ReturnResponse(c, PARAMS_INVALID, err.Error())
		return false
	}
	return true
}

func ParseQuery(c *gin.Context, obj any) bool {
	err := c.ShouldBindQuery(obj)
	if err != nil {
		ReturnResponse(c, PARAMS_INVALID, err.Error())
		return false
	}
	return true
}

func GetUserID(c *gin.Context) int64 {
	userID, exists := c.Get(USER_ID)
	if !exists {
		ReturnResponse(c, TOKEN_INVALID, "user not found")
		return 0
	}
	return userID.(int64)
}
