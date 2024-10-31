package statisticservice

import (
	"fmt"
	"gin_study/api/consts"
	"gin_study/factory"
	"gin_study/gen/query"
	"gin_study/gen/request"
	"gin_study/gen/response"
	"sync"
	"time"
)

// 定义一个结构体来接收和
type SumResult struct {
	SumSpentTime int64 `gorm:"column:sum"`
}

func GetPieValue(userID int64, req *request.GetPieValueRequest) (*[]response.PieValueResponse, error) {
	if req.StartTime > req.EndTime {
		return nil, &consts.ApiErr{Code: consts.PARAMS_INVALID, Msg: "Start time must be earlier than end time."}
	}

	respons := []response.PieValueResponse{}

	categories := req.Categories

	if len(categories) == 0 {
		err := query.Category.Select(query.Category.ID, query.Category.Name).Where(query.Category.UserID.Eq(userID)).Where(query.Category.Level.Eq(1)).Scan(&categories)
		if err != nil {
			return nil, err
		}
	}

	var wg sync.WaitGroup
	wg.Add(len(categories)) // 并发查询数据
	for _, category := range categories {
		go func(category *response.CategoriesResponse) {
			defer wg.Done() // 标记任务完成
			// 查询该分类下的任务，计算花费时间总和
			sumSpent := getSum(category.ID, userID, req.StartTime, req.EndTime)
			respons = append(respons, response.PieValueResponse{
				Value:        sumSpent,
				CategoryName: category.Name,
				CategoryID:   category.ID,
			})
		}(&category)
	}
	wg.Wait() // 所有任务完成后，主 goroutine返回
	return &respons, nil
}

func GetLineValue(userID int64, req *request.GetLineValueRequest) (*[]response.LineValueResponse, error) {
	respons := []response.LineValueResponse{}
	categories := req.Categories
	if len(categories) == 0 {
		err := query.Category.Select(query.Category.ID, query.Category.Name).Where(query.Category.UserID.Eq(userID)).Where(query.Category.Level.Eq(1)).Scan(&categories)
		if err != nil {
			return nil, err
		}
	}
	// 使用北京时区
	beijingLocation, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	wg.Add(len(categories)) // 并发查询数据
	for _, category := range categories {
		go func(category *response.CategoriesResponse) {
			defer wg.Done() // 标记任务完成
			spots := getSpot(req.TimeType, category.ID, userID, beijingLocation)
			respons = append(respons, response.LineValueResponse{
				Value:        spots,
				CategoryName: category.Name,
				CategoryID:   category.ID,
			})
		}(&category)
	}
	wg.Wait() // 所有任务完成后，主 goroutine返回
	return &respons, nil
}

func getSpot(timeType string, categoryID, userID int64, beijingLocation *time.Location) *[]response.LineSpots {
	var spots *[]response.LineSpots
	yCounts := 0
	now := time.Now().In(beijingLocation)
	switch timeType {
	case "day":
		spots = getDaySpots(now, categoryID, userID, yCounts, beijingLocation)
	case "week":
		spots = getWeekSpots(now, categoryID, userID, yCounts, beijingLocation)
	case "month":
		spots = getMonthSpots(now, categoryID, userID, yCounts, beijingLocation)
	case "year":
		spots = getYearSpots(now, categoryID, userID, yCounts, beijingLocation)
	default:
		return spots
	}
	return spots
}

func getSum(categoryID int64, userID int64, startTime, endTime int64) int64 {
	likeStr1 := fmt.Sprintf("%%,%d", categoryID)
	likeStr2 := fmt.Sprintf("%%,%d,%%", categoryID)

	whereCommon1 := query.Task.Where(query.Task.UserID.Eq(userID)).Where(query.Task.EndTime.Between(startTime, endTime))
	whereCommon2 := query.Task.Where(query.Task.UserID.Eq(userID)).Where(query.Task.EndTime.Between(startTime, endTime))

	var sumResult SumResult
	err := whereCommon1.Where(query.Task.CategoryPath.Like(likeStr1)).Or(whereCommon2.Where(query.Task.CategoryPath.Like(likeStr2))).Select(query.Task.SpentTime.Sum().As("sum")).Scan(&sumResult)
	if err != nil {
		return 0
	}
	return sumResult.SumSpentTime
}

func getDaySpots(now time.Time, categoryID int64, userID int64, yCounts int, location *time.Location) *[]response.LineSpots {
	var spots []response.LineSpots
	startTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, location)
	endTime := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, location)
	for yCounts < 12 {
		startTimeMillis := factory.GetMillis(startTime)
		endTimeMillis := factory.GetMillis(endTime)
		spendTime := getSum(categoryID, userID, startTimeMillis, endTimeMillis)
		spots = append(spots, response.LineSpots{
			X: startTimeMillis,
			Y: spendTime,
		})
		startTime = startTime.AddDate(0, 0, -1)
		endTime = endTime.AddDate(0, 0, -1)
		yCounts++
	}
	return &spots
}

func getWeekSpots(now time.Time, categoryID int64, userID int64, yCounts int, location *time.Location) *[]response.LineSpots {
	var spots []response.LineSpots
	weekday := now.Weekday()
	// 计算周一的时间
	startOfWeek := now.AddDate(0, 0, -int(weekday-1))
	endOfWeek := startOfWeek.AddDate(0, 0, 6)
	startTime := time.Date(startOfWeek.Year(), startOfWeek.Month(), startOfWeek.Day(), 0, 0, 0, 0, location)
	endTime := time.Date(endOfWeek.Year(), endOfWeek.Month(), endOfWeek.Day(), 23, 59, 59, 999999999, location)
	for yCounts < 12 {
		startTimeMillis := factory.GetMillis(startTime)
		endTimeMillis := factory.GetMillis(endTime)
		spendTime := getSum(categoryID, userID, startTimeMillis, endTimeMillis)

		spots = append(spots, response.LineSpots{
			X: startTimeMillis,
			Y: spendTime,
		})
		startTime = startTime.AddDate(0, 0, -6)
		endTime = endTime.AddDate(0, 0, -6)
		yCounts++
	}
	return &spots
}

func getMonthSpots(now time.Time, categoryID int64, userID int64, yCounts int, location *time.Location) *[]response.LineSpots {
	var spots []response.LineSpots
	startTime := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, location)
	nextMonth := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, location)
	endOfMonth := nextMonth.AddDate(0, 0, -1)
	endTime := time.Date(endOfMonth.Year(), endOfMonth.Month(), endOfMonth.Day(), 23, 59, 59, 999999999, location)
	for yCounts < 12 {
		startTimeMillis := factory.GetMillis(startTime)
		endTimeMillis := factory.GetMillis(endTime)
		spendTime := getSum(categoryID, userID, startTimeMillis, endTimeMillis)

		spots = append(spots, response.LineSpots{
			X: startTimeMillis,
			Y: spendTime,
		})
		endOfMonth = startTime.AddDate(0, 0, -1)
		endTime = time.Date(endOfMonth.Year(), endOfMonth.Month(), endOfMonth.Day(), 23, 59, 59, 999999999, location)
		startTime = time.Date(startTime.Year(), startTime.Month()-1, 1, 0, 0, 0, 0, location)
		yCounts++
	}
	return &spots
}

func getYearSpots(now time.Time, categoryID int64, userID int64, yCounts int, location *time.Location) *[]response.LineSpots {
	var spots []response.LineSpots
	startTime := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, location)
	nextYear := time.Date(now.Year()+1, 1, 1, 0, 0, 0, 0, location)
	endOfYear := nextYear.AddDate(0, 0, -1)
	endTime := time.Date(endOfYear.Year(), endOfYear.Month(), endOfYear.Day(), 23, 59, 59, 999999999, location)
	for yCounts < 12 {
		startTimeMillis := factory.GetMillis(startTime)
		endTimeMillis := factory.GetMillis(endTime)
		spendTime := getSum(categoryID, userID, startTimeMillis, endTimeMillis)
		spots = append(spots, response.LineSpots{
			X: startTimeMillis,
			Y: spendTime,
		})
		endOfYear = startTime.AddDate(0, 0, -1)
		endTime = time.Date(endOfYear.Year(), endOfYear.Month(), endOfYear.Day(), 23, 59, 59, 999999999, location)
		startTime = time.Date(startTime.Year()-1, 1, 1, 0, 0, 0, 0, location)
		yCounts++
	}
	return &spots
}
