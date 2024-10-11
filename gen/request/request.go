package request

type LoginRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetTasksRequest struct {
	Page int `form:"page" binding:"required,gte=1"`
	Size int `form:"size" binding:"required,gte=1,lte=100"`
}

type SaveTaskRequest struct {
	ID          int64  `json:"id" `
	Description string `json:"description" binding:"required"`
	CategoryID  int64  `json:"categoryID" binding:"required"`
	StartTime   int64  `json:"startTime" binding:"required"`
	EndTime     int64  `json:"endTime" binding:"required"`
}
