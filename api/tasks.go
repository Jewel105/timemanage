package api

import (
	"gin_study/gen/models"
	"gin_study/gen/query"
	"gin_study/gen/request"
	"gin_study/gen/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskApi struct{}

func (t TaskApi) GetList(c *gin.Context) {
	userID := GetUserID(c)
	if userID == 0 {
		return
	}
	req := request.GetTasksRequest{}
	if !ParseQuery(c, &req) {
		return
	}
	offset := (req.Page - 1) * req.Size
	tasks, err := query.Task.Where(query.Task.UserID.Eq(userID)).Limit(req.Size).Offset(offset).Find()
	if err != nil {
		ReturnResponse(c, SYSTEM_ERROR, err.Error())
		return
	}
	res := response.PageResponse{
		Page: req.Page,
		Size: req.Size,
		Data: tasks,
	}
	DealResponse(c, res, err)
}

func (t TaskApi) SaveTask(c *gin.Context) {
	userID := GetUserID(c)
	if userID == 0 {
		return
	}
	req := request.SaveTaskRequest{}
	if !ParseJson(c, &req) {
		return
	}

	// Category not found.
	_, errQuery := query.Category.Where(query.Category.UserID.Eq(userID)).Where(query.Category.ID.Eq(req.CategoryID)).First()

	if errQuery != nil {
		ReturnResponse(c, CLIENT_ERROR, "Category not found.")
		return
	}

	task := models.Task{
		ID:          req.ID,
		UserID:      userID,
		Description: req.Description,
		CategoryID:  req.CategoryID,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
	}
	task.SpentTime = task.EndTime - task.StartTime

	tx := query.Q.Begin()
	err := query.Task.Save(&task)
	if err != nil {
		err = tx.Rollback()
	} else {
		err = tx.Commit()
	}
	DealResponse(c, task.ID, err)
}

func (t TaskApi) DeleteTask(c *gin.Context) {
	userID := GetUserID(c)
	if userID == 0 {
		return
	}
	idStr := c.Param("id")
	id, parseErr := strconv.ParseInt(idStr, 10, 64)
	if parseErr != nil {
		ReturnResponse(c, SYSTEM_ERROR, parseErr.Error())
		return
	}

	tx := query.Q.Begin()
	qIDO := query.Task.Where(query.Task.ID.Eq(id)).Where(query.Task.UserID.Eq(userID))
	qIDO.First()
	_, err := qIDO.Delete()
	if err != nil {
		DealResponse(c, nil, err)
		tx.Rollback()
		return
	}
	tx.Commit()
	DealResponse(c, nil, err)
}
