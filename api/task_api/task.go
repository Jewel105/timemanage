package taskapi

import (
	"gin_study/api"
	"gin_study/gen/request"
	taskservice "gin_study/service/task_service"

	"github.com/gin-gonic/gin"
)

// @Id TaskGetList
// @Summary 查询任务列表
// @Description 查询任务列表
// @Tags 任务API
// @Accept  json
// @Produce application/json
// @Param token header string false "enjmcvhdwernxhcuvyudfdjfhkjxkjaoerpq"
// @Param page query int false "1"
// @Param size query int false "10"
// @Param startTime query int true "开始时间"
// @Param endTime query int true  "结束时间"
// @success 200 {object} []response.TasksResponse "success"
// @Router /tasks/list [get]
func GetList(c *gin.Context) {
	userID := api.GetUserID(c)
	if userID == 0 {
		return
	}
	req := request.GetTasksRequest{}
	if !api.ParseQuery(c, &req) {
		return
	}
	res, err := taskservice.GetList(userID, &req)
	api.DealResponse(c, res, err)
}

// @Id SaveTask
// @Summary 保存或修改任务
// @Description 保存或修改任务
// @Tags 任务API
// @Accept  json
// @Produce application/json
// @Param token header string false "enjmcvhdwernxhcuvyudfdjfhpq"
// @Param req body request.SaveTaskRequest true "Json"
// @success 200 int64 taskID "success"
// @Router  /tasks/save [post]
func SaveTask(c *gin.Context) {
	userID := api.GetUserID(c)
	if userID == 0 {
		return
	}
	req := request.SaveTaskRequest{}
	if !api.ParseJson(c, &req) {
		return
	}
	taskID, err := taskservice.SaveTask(userID, &req)
	api.DealResponse(c, taskID, err)
}

// @Id DeleteTask
// @Summary 删除任务
// @Description 删除任务
// @Tags 任务API
// @Accept  json
// @Produce application/json
// @Param token header string false "enjmcvhdwernxhcuvyudfdjfhpq"
// @Param id path int true "任务ID"
// @success 200 boolean ture "success"
// @Router  /tasks/delete/:id [post]
func DeleteTask(c *gin.Context) {
	userID := api.GetUserID(c)
	if userID == 0 {
		return
	}
	idStr := c.Param("id")
	err := taskservice.DeleteTask(userID, idStr)
	api.DealResponse(c, true, err)
}

// @Id GetLastEndTime
// @Summary 获取最后一个结束时间
// @Description 获取最后一个结束时间
// @Tags 任务API
// @Accept  json
// @Produce application/json
// @Param token header string false "enjmcvhdwernxhcuvyudfdjfhkjxkjaoerpq"
// @success 200 {integer} int64 "success"
// @Router /tasks/last/time [get]
func GetLastEndTime(c *gin.Context) {
	userID := api.GetUserID(c)
	if userID == 0 {
		return
	}
	res, err := taskservice.GetLastEndTime(userID)
	api.DealResponse(c, res, err)
}
