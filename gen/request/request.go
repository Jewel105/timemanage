package request

type LoginRequest struct {
	Name     string `json:"name" binding:"required"`     // 用户名或邮箱
	Password string `json:"password" binding:"required"` // 密码
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`                                // 用户名
	Password string `json:"password" binding:"required"`                            // 新密码
	Email    string `json:"email" example:"test@mail.com" binding:"required,email"` // 邮箱
	Code     string `json:"code" example:"888888" binding:"required,len=6"`         // 验证码
}

type GetTasksRequest struct {
	Page      int `form:"page,default=1" binding:"gte=1"`          // 页码
	Size      int `form:"size,default=10" binding:"gte=1,lte=100"` // 每页数量
	StartTime int `form:"startTime" binding:"required"`            // 开始时间
	EndTime   int `form:"endTime" binding:"required"`              // 开始时间
}

type SaveTaskRequest struct {
	ID          int64  `json:"id"`                            // 任务ID
	Description string `json:"description"`                   // 任务描述
	CategoryID  int64  `json:"categoryID" binding:"required"` // 任务所属分类ID
	StartTime   int64  `json:"startTime" binding:"required"`  // 任务开始时间
	EndTime     int64  `json:"endTime" binding:"required"`    // 任务结束时间
}

type GetCategoriesRequest struct {
	ParentID int64 `form:"parentID,default=0"` // 上级分类ID，默认0
}

type SaveCategoryRequest struct {
	ID       int64  `json:"id"`                      // 分类ID
	Name     string `json:"name" binding:"required"` // 分类名称
	ParentID int64  `json:"parentID"`                // 上级分类ID
}

type SendCodeRequest struct {
	Email string `json:"email" example:"test@mail.com" binding:"required,email"` // 邮箱
}

type RegisterEquipmentRequest struct {
	Vender           string `json:"vender"`           // 供应商
	Type             string `json:"type"`             // 设备类型
	Sn               string `json:"sn"`               // 序列号
	Imei1            string `json:"imei1"`            // IMEI1
	Imei0            string `json:"imei0"`            // IMEI0
	Os               string `json:"os"`               // 所属操作系统
	IsPhysicalDevice int    `json:"isPhysicalDevice"` // 是否为物理设备
}

type LogErrorRequest struct {
	Version string `json:"version"` // 版本
	Stack   string `json:"stack"`   // 堆栈
	Error   string `json:"error"`   // 错误信息
}
