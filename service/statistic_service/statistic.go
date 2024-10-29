package statisticservice

import (
	"fmt"
	"gin_study/gen/query"
	"gin_study/gen/request"
	"gin_study/gen/response"
)

// 定义一个结构体来接收和
type SumResult struct {
	SumSpentTime int64 `gorm:"column:sum"`
}

func GetPieValue(userID int64, req *request.GetPieValueRequest) (*[]response.PieValueResponse, error) {
	respons := []response.PieValueResponse{}
	categories := req.Categories

	if len(categories) == 0 {
		err := query.Category.Select(query.Category.ID, query.Category.Name).Where(query.Category.UserID.Eq(userID)).Where(query.Category.Level.Eq(1)).Scan(&categories)
		if err != nil {
			return nil, err
		}
	}

	for _, category := range categories {
		c := make(chan int64)
		// 查询该分类下的任务，计算花费时间总和
		go func() {
			sumSpent, err := getSum(category.ID, userID, req)
			if err != nil {
				c <- 0
			} else {
				c <- sumSpent
			}
		}()

		sumSpent := <-c
		respons = append(respons, response.PieValueResponse{
			Value:        sumSpent,
			CategoryName: category.Name,
			CategoryID:   category.ID,
		})
	}
	return &respons, nil
}

func getSum(categoryID int64, userID int64, req *request.GetPieValueRequest) (int64, error) {
	likeStr1 := fmt.Sprintf("%%,%d", categoryID)
	likeStr2 := fmt.Sprintf("%%,%d,%%", categoryID)

	whereCommon1 := query.Task.Where(query.Task.UserID.Eq(userID)).Where(query.Task.EndTime.Between(req.StartTime, req.EndTime))
	whereCommon2 := query.Task.Where(query.Task.UserID.Eq(userID)).Where(query.Task.EndTime.Between(req.StartTime, req.EndTime))

	var sumResult SumResult
	err := whereCommon1.Where(query.Task.CategoryPath.Like(likeStr1)).Or(whereCommon2.Where(query.Task.CategoryPath.Like(likeStr2))).Select(query.Task.SpentTime.Sum().As("sum")).Scan(&sumResult)
	if err != nil {
		return 0, err
	}
	return sumResult.SumSpentTime, nil
}
