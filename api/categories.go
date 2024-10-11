package api

import (
	"fmt"
	"gin_study/gen/models"
	"gin_study/gen/query"
	"gin_study/gen/request"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type CategoryApi struct{}

func (t CategoryApi) GetList(c *gin.Context) {
	userID := GetUserID(c)
	if userID == 0 {
		return
	}
	req := request.GetCategoriesRequest{}
	if !ParseQuery(c, &req) {
		return
	}

	categories, err := query.Category.Where(query.Category.UserID.Eq(userID)).Where(
		query.Category.ParentID.Eq(req.ParentID)).Find()

	DealResponse(c, categories, err)
}

func (t CategoryApi) SaveCategory(c *gin.Context) {
	userID := GetUserID(c)
	if userID == 0 {
		return
	}
	req := request.SaveCategoryRequest{}
	if !ParseJson(c, &req) {
		return
	}
	qUserCategory := query.Category.Where(query.Category.UserID.Eq(userID))

	category := models.Category{
		ID:       req.ID,
		UserID:   userID,
		ParentID: req.ParentID,
		Name:     req.Name,
	}

	parentCategory, errQuery := qUserCategory.Where(query.Category.ID.Eq(req.ParentID)).First()
	// 不存在上级分类，则保存为根分类
	if errQuery != nil {
		category.Path = "0"
		category.ParentID = 0
	} else {
		category.Path = fmt.Sprintf("%s,%d", parentCategory.Path, category.ParentID)
	}
	category.Level = len(strings.Split(category.Path, ","))

	tx := query.Q.Begin()
	err := query.Category.Save(&category)
	if err != nil {
		err = tx.Rollback()
	} else {
		err = tx.Commit()
	}
	DealResponse(c, category.ID, err)
}

func (t CategoryApi) DeleteCategory(c *gin.Context) {
	userID := GetUserID(c)
	if userID == 0 {
		return
	}
	idStr := c.Param("id")
	id, parseErr := strconv.ParseInt(idStr, 10, 64)
	if parseErr != nil {
		ReturnResponse(c, SYSTEM_ERROR, parseErr.Error())
		return
	}
	tx := query.Q.Begin()
	qIDO := query.Category.Where(query.Category.ID.Eq(id)).Where(query.Category.UserID.Eq(userID))
	_, err := qIDO.Delete()

	if err != nil {
		if err := tx.Rollback(); err != nil {
			return
		}
	} else {
		if err := tx.Commit(); err != nil {
			return
		}
	}
	DealResponse(c, true, err)
}
