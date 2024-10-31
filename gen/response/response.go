package response

type TasksResponse struct {
	ID          int64  `json:"id"`          // 任务ID
	Description string `json:"description"` // 任务描述
	SpentTime   int64  `json:"spentTime"`   // 花费时间
	Categories  string `json:"categories"`  // 任务所属分类
	CategoryID  int64  `json:"categoryID"`  // 任务所属分类ID
	StartTime   int64  `json:"startTime"`   // 任务开始时间
	EndTime     int64  `json:"endTime"`     // 任务结束时间
}

type CategoriesResponse struct {
	ID       int64  `json:"id"`       // 分类ID
	Name     string `json:"name"`     // 分类名称
	ParentID int64  `json:"parentID"` // 上级分类ID
	UserID   int64  `json:"userID"`   // 创建该分类的用户ID
	Path     string `json:"path"`     // 分类路径
	Level    int    `json:"level"`    // 分类等级
}

type PieValueResponse struct {
	CategoryID   int64  `json:"categoryID"`   // 所需查询的分类ID
	CategoryName string `json:"categoryName"` // 所需查询的分类名称
	Value        int64  `json:"value"`        // 值
}

type LineValueResponse struct {
	CategoryID   int64        `json:"categoryID"`   // 所需查询的分类ID
	CategoryName string       `json:"categoryName"` // 所需查询的分类名称
	Value        *[]LineSpots `json:"value"`        // 值
}

type LineSpots struct {
	X int64 `json:"x"` // x轴值
	Y int64 `json:"y"` // y轴值
}
