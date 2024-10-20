package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin_study/api/consts"
	"gin_study/factory"
	"gin_study/logger"
	"io"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

type TokenHeader struct {
	Token string `json:"token" header:"token" binding:"required" err:"token error"`
}

func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(zap.Any("gin recover error", err))
			ReturnResponse(c, consts.SYSTEM_ERROR, fmt.Sprintf("%v", err))
			//终止后续接口调用，不加的话recover异常之后，还会继续执行后续代码
			c.Abort()
		}
	}()
	c.Next()
}

func VerifyToken(c *gin.Context) {
	var token TokenHeader
	if err := c.ShouldBindHeader(&token); err != nil {
		ReturnResponse(c, consts.TOKEN_INVALID, err.Error())
		c.Abort()
		return
	}
	userClaims, err := factory.DecodeToken(token.Token)
	if err != nil {
		ReturnResponse(c, consts.TOKEN_INVALID, err.Error())
		c.Abort()
		return
	}

	c.Set(consts.USER_ID, userClaims.UserID)
	c.Next()
}

func SaveEquipmentID(c *gin.Context) {
	equipmentIDStr := c.GetHeader("Equipment")
	if equipmentIDStr == "" {
		c.Next()
		return
	}
	equipmentID, err := strconv.ParseInt(equipmentIDStr, 10, 64)
	if err != nil {
		c.Next()
		return
	}
	c.Set(consts.EQUIPMENT_ID, equipmentID)
	c.Next()
}

func RecordLog(c *gin.Context) {
	// 开始时间
	startTime := time.Now()
	bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = bodyLogWriter
	// 打印请求信息
	reqBody, _ := c.GetRawData()
	// 使后续请求能收到请求参数
	c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
	//处理请求
	c.Next()

	// //结束时间
	endTime := time.Now()
	duration := endTime.Sub(startTime).Milliseconds()
	// 尝试解析 response body 为 JSON
	var responseBody interface{}
	responseBytes := bodyLogWriter.body.Bytes()
	err := json.Unmarshal(responseBytes, &responseBody)
	if err != nil {
		// 如果解析失败，就作为字符串输出
		responseBody = string(responseBytes)

	}
	// 记录gin-log
	logger.Info("gin-log",
		zap.String("Client IP", c.ClientIP()),
		zap.Int("Code", c.Writer.Status()),
		zap.String("Method", c.Request.Method),
		zap.String("URI", c.Request.RequestURI),
		zap.String("Protocol", c.Request.Proto),
		zap.String("Token", c.GetHeader("token")),
		zap.String("UserAgent", c.Request.UserAgent()),
		zap.Any("StartTime", startTime),
		zap.Any("EndTime", endTime),
		zap.Any("Duration", duration),
		zap.String("RequestBody", string(reqBody)),
		zap.Any("ResponseBody", responseBody),
	)
}
