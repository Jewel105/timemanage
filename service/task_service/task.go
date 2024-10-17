package taskservice

import (
	"fmt"
	"gin_study/api/consts"
	"gin_study/gen/models"
	"gin_study/gen/mysql"
	"gin_study/gen/query"
	"gin_study/gen/request"
	"gin_study/gen/response"
	"strconv"
)

func GetList(userID int64, req *request.GetTasksRequest) (*response.PageResponse, error) {
	offset := (req.Page - 1) * req.Size
	tasks := []response.TasksRespose{}
	sql := fmt.Sprintf(
		`
		SELECT 
			tasks.*,
			GROUP_CONCAT(categories.name ORDER BY categories.id) AS categories
		FROM 
				tasks 
		LEFT JOIN 
				categories ON FIND_IN_SET(categories.id, tasks.category_path) 
		WHERE 
				tasks.user_id = %d AND tasks.delete_time IS NULL
		GROUP BY 
				tasks.id
		LIMIT %d
		OFFSET %d
			`, userID, req.Size, offset,
	)

	query.Task.UnderlyingDB().Raw(sql).Scan(&tasks)

	res := response.PageResponse{
		Page: req.Page,
		Size: req.Size,
		Data: tasks,
	}
	return &res, nil
}

func SaveTask(userID int64, req *request.SaveTaskRequest) (int64, error) {
	// Category not found.
	category, err := query.Category.Where(query.Category.UserID.Eq(userID)).Where(query.Category.ID.Eq(req.CategoryID)).First()
	if err != nil {
		return 0, &consts.ApiErr{Code: consts.BAD_REQUEST, Msg: "Category not found."}
	}

	task := models.Task{
		ID:           req.ID,
		UserID:       userID,
		Description:  req.Description,
		CategoryPath: fmt.Sprintf("%s,%d", category.Path, req.CategoryID),
		CategoryID:   req.CategoryID,
		StartTime:    req.StartTime,
		EndTime:      req.EndTime,
	}

	task.SpentTime = task.EndTime - task.StartTime

	tx := query.Q.Begin()
	err = query.Task.Save(&task)
	err = mysql.DeferTx(tx, err)
	return task.ID, err
}

func DeleteTask(userID int64, idStr string) error {
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return err
	}

	tx := query.Q.Begin()
	taskQuery := query.Task.Where(query.Task.ID.Eq(id)).Where(query.Task.UserID.Eq(userID))

	count, err := taskQuery.Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return &consts.ApiErr{Code: consts.BAD_REQUEST, Msg: "Task not exists."}
	}
	_, err = taskQuery.Delete()
	err = mysql.DeferTx(tx, err)
	return err
}
