package request

type LoginRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SaveTaskRequest struct {
	ID          uint    `json:"id" `
	Description string  `json:"description" binding:"required"`
	SpentTime   float64 `json:"spentTime" binding:"required"`
	CategoryID  int64   `json:"categoryID" binding:"required"`
	StartTime   int64   ` json:"startTime" binding:"required"`
	EndTime     int64   ` json:"endTime" binding:"required"`
}
