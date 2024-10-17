package request

type LoginRequest struct {
	Name     string `json:"name" binding:"required"`     // 用户名
	Password string `json:"password" binding:"required"` // 密码
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`                                        // 用户名
	Password string `json:"password" binding:"required"`                                    // 密码
	Email    string `json:"email" example:"test@mail.com" binding:"required,email,max=255"` // 邮箱
	Code     string `json:"code" example:"888888" binding:"required,len=6"`                 // 验证码
}

type GetTasksRequest struct {
	Page int `form:"page,default=1" binding:"gte=1"`          // 页码
	Size int `form:"size,default=10" binding:"gte=1,lte=100"` // 每页数量
}

type SaveTaskRequest struct {
	ID          int64  `json:"id"`
	Description string `json:"description" binding:"required"` // 任务描述
	CategoryID  int64  `json:"categoryID" binding:"required"`  // 任务所属分类ID
	StartTime   int64  `json:"startTime" binding:"required"`   // 任务开始时间
	EndTime     int64  `json:"endTime" binding:"required"`     // 任务结束时间
}

type GetCategoriesRequest struct {
	ParentID int64 `form:"parentID,default=0"`
}

type SaveCategoryRequest struct {
	ID       int64  `json:"id"`
	Name     string `json:"name" binding:"required"` // 分类名称
	ParentID int64  `json:"parentID"`                // 上级分类ID
}

type SendCodeRequest struct {
	Email string `json:"email" example:"test@mail.com" binding:"required,email,max=255"` // 邮箱
}
