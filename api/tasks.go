package api

import (
	"gin_study/gen/models"
	"gin_study/gen/query"
	"gin_study/gen/request"

	"github.com/gin-gonic/gin"
)

type TaskController struct{}

func (t TaskController) GetList(c *gin.Context) {
	userID := GetUserID(c)
	if userID == 0 {
		return
	}
	tasks, err := query.Task.Where(query.Task.UserID.Eq(userID)).Find()
	DealResponse(c, tasks, err)
}

func (t TaskController) SaveTask(c *gin.Context) {
	userID := GetUserID(c)
	if userID == 0 {
		return
	}
	req := request.SaveTaskRequest{}
	if !ParseJson(c, &req) {
		return
	}
	task := models.Task{
		UserID:      userID,
		Description: req.Description,
		SpentTime:   req.SpentTime,
		CategoryID:  req.CategoryID,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
	}
	task.ID = req.ID

	tx := query.Q.Begin()
	err := query.Task.Save(&task)
	if err != nil {
		err = tx.Rollback()
	} else {
		err = tx.Commit()
	}
	DealResponse(c, task.ID, err)
}
