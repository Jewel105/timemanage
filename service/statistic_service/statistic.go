package statisticservice

import (
	"fmt"
	"gin_study/factory"
	"gin_study/gen/query"
	"gin_study/gen/request"
	"gin_study/gen/response"
	"time"
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
		// 查询该分类下的任务，计算花费时间总和
		sumSpent, err := getSum(category.ID, userID, req.StartTime, req.EndTime)
		if err != nil {
			sumSpent = 0
		}
		respons = append(respons, response.PieValueResponse{
			Value:        sumSpent,
			CategoryName: category.Name,
			CategoryID:   category.ID,
		})
	}
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

	for _, category := range categories {
		var spots *[]response.LineSpots
		yCounts := 0
		now := time.Now().In(beijingLocation)
		switch req.TimeType {
		case "day":
			spots, err = getDaySpots(now, category.ID, userID, yCounts, beijingLocation)
			if err != nil {
				return nil, err
			}
		case "week":
			spots, err = getWeekSpots(now, category.ID, userID, yCounts, beijingLocation)
			if err != nil {
				return nil, err
			}
		case "month":
			spots, err = getMonthSpots(now, category.ID, userID, yCounts, beijingLocation)
			if err != nil {
				return nil, err
			}
		case "year":
			spots, err = getYearSpots(now, category.ID, userID, yCounts, beijingLocation)
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("invalid time type: %s", req.TimeType)
		}

		respons = append(respons, response.LineValueResponse{
			Value:        spots,
			CategoryName: category.Name,
			CategoryID:   category.ID,
		})
	}
	return &respons, nil
}

func getSum(categoryID int64, userID int64, startTime, endTime int64) (int64, error) {
	likeStr1 := fmt.Sprintf("%%,%d", categoryID)
	likeStr2 := fmt.Sprintf("%%,%d,%%", categoryID)

	whereCommon1 := query.Task.Where(query.Task.UserID.Eq(userID)).Where(query.Task.EndTime.Between(startTime, endTime))
	whereCommon2 := query.Task.Where(query.Task.UserID.Eq(userID)).Where(query.Task.EndTime.Between(startTime, endTime))

	var sumResult SumResult
	err := whereCommon1.Where(query.Task.CategoryPath.Like(likeStr1)).Or(whereCommon2.Where(query.Task.CategoryPath.Like(likeStr2))).Select(query.Task.SpentTime.Sum().As("sum")).Scan(&sumResult)
	if err != nil {
		return 0, err
	}
	return sumResult.SumSpentTime, nil
}

func getDaySpots(now time.Time, categoryID int64, userID int64, yCounts int, location *time.Location) (*[]response.LineSpots, error) {
	var spots []response.LineSpots
	startTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, location)
	endTime := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, location)
	for yCounts < 12 {
		startTimeMillis := factory.GetMillis(startTime)
		endTimeMillis := factory.GetMillis(endTime)
		spendTime, err := getSum(categoryID, userID, startTimeMillis, endTimeMillis)
		if err != nil {
			return nil, err
		}
		spots = append(spots, response.LineSpots{
			X: startTimeMillis,
			Y: spendTime,
		})
		startTime = startTime.AddDate(0, 0, -1)
		endTime = endTime.AddDate(0, 0, -1)
		yCounts++
	}
	return &spots, nil
}

func getWeekSpots(now time.Time, categoryID int64, userID int64, yCounts int, location *time.Location) (*[]response.LineSpots, error) {
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
		spendTime, err := getSum(categoryID, userID, startTimeMillis, endTimeMillis)
		if err != nil {
			return nil, err
		}
		spots = append(spots, response.LineSpots{
			X: startTimeMillis,
			Y: spendTime,
		})
		startTime = startTime.AddDate(0, 0, -6)
		endTime = endTime.AddDate(0, 0, -6)
		yCounts++
	}
	return &spots, nil
}

func getMonthSpots(now time.Time, categoryID int64, userID int64, yCounts int, location *time.Location) (*[]response.LineSpots, error) {
	var spots []response.LineSpots
	startTime := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, location)
	nextMonth := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, location)
	endOfMonth := nextMonth.AddDate(0, 0, -1)
	endTime := time.Date(endOfMonth.Year(), endOfMonth.Month(), endOfMonth.Day(), 23, 59, 59, 999999999, location)
	for yCounts < 12 {
		startTimeMillis := factory.GetMillis(startTime)
		endTimeMillis := factory.GetMillis(endTime)
		spendTime, err := getSum(categoryID, userID, startTimeMillis, endTimeMillis)
		if err != nil {
			return nil, err
		}
		spots = append(spots, response.LineSpots{
			X: startTimeMillis,
			Y: spendTime,
		})
		endOfMonth = startTime.AddDate(0, 0, -1)
		endTime = time.Date(endOfMonth.Year(), endOfMonth.Month(), endOfMonth.Day(), 23, 59, 59, 999999999, location)
		startTime = time.Date(startTime.Year(), startTime.Month()-1, 1, 0, 0, 0, 0, location)
		yCounts++
	}
	return &spots, nil
}

func getYearSpots(now time.Time, categoryID int64, userID int64, yCounts int, location *time.Location) (*[]response.LineSpots, error) {
	var spots []response.LineSpots
	startTime := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, location)
	nextYear := time.Date(now.Year()+1, 1, 1, 0, 0, 0, 0, location)
	endOfYear := nextYear.AddDate(0, 0, -1)
	endTime := time.Date(endOfYear.Year(), endOfYear.Month(), endOfYear.Day(), 23, 59, 59, 999999999, location)
	for yCounts < 12 {
		startTimeMillis := factory.GetMillis(startTime)
		endTimeMillis := factory.GetMillis(endTime)
		spendTime, err := getSum(categoryID, userID, startTimeMillis, endTimeMillis)
		if err != nil {
			return nil, err
		}
		spots = append(spots, response.LineSpots{
			X: startTimeMillis,
			Y: spendTime,
		})
		endOfYear = startTime.AddDate(0, 0, -1)
		endTime = time.Date(endOfYear.Year(), endOfYear.Month(), endOfYear.Day(), 23, 59, 59, 999999999, location)
		startTime = time.Date(startTime.Year()-1, 1, 1, 0, 0, 0, 0, location)
		yCounts++
	}
	return &spots, nil
}
