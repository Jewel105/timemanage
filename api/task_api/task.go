package taskapi

import (
	"gin_study/api"
	"gin_study/gen/request"
	taskservice "gin_study/service/task_service"

	"github.com/gin-gonic/gin"
)

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

func DeleteTask(c *gin.Context) {
	userID := api.GetUserID(c)
	if userID == 0 {
		return
	}
	idStr := c.Param("id")
	err := taskservice.DeleteTask(userID, idStr)
	api.DealResponse(c, true, err)
}
