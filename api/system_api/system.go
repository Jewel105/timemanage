package systemapi

import (
	"gin_study/api"
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
// @success 200 boolean ture "success"
// @Router  /register/equipment [post]
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
// @success 200 boolean ture "success"
// @Router  /log/error [post]
func LogError(c *gin.Context) {
	userID := api.GetUserID(c)
	if userID == 0 {
		return
	}
	req := request.LogErrorRequest{}
	if !api.ParseJson(c, &req) {
		return
	}
	// equipmentID, err := systemservice.RegisterEquipment(&req)
	// api.DealResponse(c, equipmentID, err)
}
