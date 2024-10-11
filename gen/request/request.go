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
	Page int `form:"page,default=1" binding:"gte=1"`
	Size int `form:"size,default=10" binding:"gte=1,lte=100"`
}

type SaveTaskRequest struct {
	ID          int64  `json:"id"`
	Description string `json:"description" binding:"required"`
	CategoryID  int64  `json:"categoryID" binding:"required"`
	StartTime   int64  `json:"startTime" binding:"required"`
	EndTime     int64  `json:"endTime" binding:"required"`
}

type GetCategoriesRequest struct {
	ParentID int64 `form:"parentID,default=0"`
}

type SaveCategoryRequest struct {
	ID       int64  `json:"id"`
	Name     string `json:"name" binding:"required"`
	ParentID int64  `json:"parentID"` // 上级分类ID
}
