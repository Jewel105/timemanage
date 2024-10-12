package categoryapi

import (
	"fmt"
	"gin_study/api"
	"gin_study/api/consts"
	"gin_study/gen/models"
	"gin_study/gen/query"
	"gin_study/gen/request"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetList(c *gin.Context) {
	userID := api.GetUserID(c)
	if userID == 0 {
		return
	}
	req := request.GetCategoriesRequest{}
	if !api.ParseQuery(c, &req) {
		return
	}

	categories, err := query.Category.Where(query.Category.UserID.Eq(userID)).Where(
		query.Category.ParentID.Eq(req.ParentID)).Find()

	api.DealResponse(c, categories, err)
}

func SaveCategory(c *gin.Context) {
	userID := api.GetUserID(c)
	if userID == 0 {
		return
	}
	req := request.SaveCategoryRequest{}
	if !api.ParseJson(c, &req) {
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
	api.DealResponse(c, category.ID, err)
}

func DeleteCategory(c *gin.Context) {
	userID := api.GetUserID(c)
	if userID == 0 {
		return
	}
	idStr := c.Param("id")
	id, parseErr := strconv.ParseInt(idStr, 10, 64)
	if parseErr != nil {
		api.ReturnResponse(c, consts.SYSTEM_ERROR, parseErr.Error())
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
	api.DealResponse(c, true, err)
}
