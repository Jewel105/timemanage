package api

import (
	"fmt"
	"gin_study/factory"
	"io"
	"os"
	"path"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
)

type TokenHeader struct {
	Token string `json:"token" header:"token" binding:"required" err:"token error"`
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

func VerifyToken(c *gin.Context) {
	var token TokenHeader
	if err := c.ShouldBindHeader(&token); err != nil {
		ReturnResponse(c, TOKEN_INVALID, err.Error())
		c.Abort()
		return
	}
	userClaims, err := factory.DecodeToken(token.Token)
	if err != nil {
		ReturnResponse(c, TOKEN_INVALID, err.Error())
		c.Abort()
		return
	}

	c.Set(USER_ID, userClaims.UserID)
	c.Next()
}
