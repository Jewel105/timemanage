package api

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type JsonStruct struct {
	Code    string      `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

const (
	SUCCESS             = "200"
	SYSTEM_ERROR        = "500"
	CLIENT_ERROR        = "400"
	SERVICE_UNAVAILABLE = "503"

	PARAMS_INVALID             = "300"
	TOKEN_INVALID              = "301"
	PROCESSING                 = "302"
	EQUIPMENT_INVALID          = "303"
	MATCH_ORDER_ERROR          = "304"
	ORDER_EXISTS               = "305"
	ACCOUNT_BALANCE_NOT_ENOUGH = "306"
	LOGIN_FAILED               = "-1"
	NOT_SUPPORTED              = "307"
	UNABLE_TO_VERIFY           = "308"
)

var MsgFlags = map[string]string{
	SUCCESS:             "SUCCESS",
	SYSTEM_ERROR:        "SYSTEM_ERROR",
	CLIENT_ERROR:        "CLIENT_ERROR",
	SERVICE_UNAVAILABLE: "SERVICE_UNAVAILABLE",

	PARAMS_INVALID:             "PARAMS_INVALID",
	TOKEN_INVALID:              "TOKEN_INVALID",
	PROCESSING:                 "PROCESSING",
	LOGIN_FAILED:               "LOGIN_FAILED",
	EQUIPMENT_INVALID:          "EQUIPMENT_INVALID",
	MATCH_ORDER_ERROR:          "MATCH_ORDER_ERROR",
	ORDER_EXISTS:               "ORDER_EXISTS",
	ACCOUNT_BALANCE_NOT_ENOUGH: "ACCOUNT_BALANCE_NOT_ENOUGH",
	NOT_SUPPORTED:              "NOT_SUPPORTED",
	UNABLE_TO_VERIFY:           "UNABLE_TO_VERIFY",
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
		ReturnResponse(c, SYSTEM_ERROR, err.Error())
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

func LoggerToFile() gin.LoggerConfig {
	if _, err := os.Stat("./runtime/log"); os.IsNotExist(err) {
		err = os.MkdirAll("./runtime/log", 0777)
		if err != nil {
			panic(fmt.Errorf("create log dir '%s' error:%s", "./runtime/log", err))
		}
	}

	timeStr := time.Now().Format("2006-01-02")
	fileName := path.Join("./runtime/log", "success_"+timeStr+".log")

	os.Stderr, _ = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	var conf = gin.LoggerConfig{
		Formatter: func(params gin.LogFormatterParams) string {
			return fmt.Sprintf("%s - %s \"%s %s %s %d \"%s\" %s\"%s\"\n",
				params.TimeStamp.Format("2006-01-02 15:04:06"),
				params.ClientIP,
				params.Method,
				params.Path,
				params.Request.Proto,
				params.StatusCode,
				params.Latency,
				params.Request.UserAgent(),
				params.ErrorMessage,
			)
		},
		Output: io.MultiWriter(os.Stdout, os.Stderr),
	}

	return conf

}

func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			if _, errDir := os.Stat("./runtime/log"); os.IsNotExist(errDir) {
				err = os.MkdirAll("./runtime/log", 0777)
				if err != nil {
					panic(fmt.Errorf("create log dir '%s' error:%s", "./runtime/log", err))
				}
			}

			timeStr := time.Now().Format("2006-01-02")
			fileName := path.Join("./runtime/log", "err_"+timeStr+".log")

			f, errFile := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

			if errFile != nil {
				fmt.Println(errFile)
			}

			timeFileStr := time.Now().Format("2006-01-02 15:04:05")
			f.WriteString("panic error time:" + timeFileStr + "\n")
			f.WriteString(fmt.Sprintf("%v", err) + "\n")
			f.WriteString("stacktrace from panic:" + string(debug.Stack()) + "\n")
			f.Close()
			ReturnResponse(c, SYSTEM_ERROR, fmt.Sprintf("%v", err))
			//终止后续接口调用，不加的话recover异常之后，还会继续执行后续代码
			c.Abort()
		}
	}()
	c.Next()
}

func ParseJson(c *gin.Context, obj any) bool {
	err := c.ShouldBindJSON(obj)
	if err != nil {
		ReturnResponse(c, PARAMS_INVALID, err.Error())
		return false
	}
	return true
}
