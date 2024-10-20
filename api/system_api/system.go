package systemapi

import (
	"gin_study/api"
	"gin_study/api/consts"
	"gin_study/gen/request"
	systemservice "gin_study/service/system_service"

	"github.com/gin-gonic/gin"
)

// @Id RegisterEquipment
// @Summary 注册设备
// @Description 注册设备
// @Tags COMMON API
// @Accept  json
// @Produce application/json
// @Param req body request.RegisterEquipmentRequest true "Json"
// @Success 200 {integer} integer true "success"
// @Router  /common/system/register/equipment [post]
func RegisterEquipment(c *gin.Context) {
	req := request.RegisterEquipmentRequest{}
	if !api.ParseJson(c, &req) {
		return
	}
	equipmentID, err := systemservice.RegisterEquipment(&req)
	api.DealResponse(c, equipmentID, err)
}

// @Id LogError
// @Summary 记录前端日志错误
// @Description 记录前端日志错误
// @Tags COMMON API
// @Accept  json
// @Produce application/json
// @Param req body request.LogErrorRequest true "Json"
// @Success 200 {integer} integer true "success"
// @Router  /common/system/log/error [post]
func LogError(c *gin.Context) {
	userID := c.GetInt64(consts.USER_ID)
	equipmentID := c.GetInt64(consts.EQUIPMENT_ID)

	req := request.LogErrorRequest{}
	if !api.ParseJson(c, &req) {
		return
	}
	logID, err := systemservice.LogError(userID, equipmentID, &req)
	api.DealResponse(c, logID, err)
}
