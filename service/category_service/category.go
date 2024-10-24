package categoryservice

import (
	"fmt"
	"gin_study/api/consts"
	"gin_study/gen/models"
	"gin_study/gen/mysql"
	"gin_study/gen/query"
	"gin_study/gen/request"
	"gin_study/gen/response"
	"strconv"
	"strings"
)

func GetList(userID int64, req *request.GetCategoriesRequest) (*[]response.CategoriesResponse, error) {
	category := []response.CategoriesResponse{}
	query.Category.Where(query.Category.UserID.Eq(userID)).Where(query.Category.ParentID.Eq(req.ParentID)).Scan(&category)
	return &category, nil
}

func SaveCategory(userID int64, req *request.SaveCategoryRequest) (int64, error) {
	categoryQuery := query.Category.Where(query.Category.UserID.Eq(userID))
	parentCategory, errQuery := categoryQuery.Where(query.Category.ID.Eq(req.ParentID)).First()

	category := models.Category{
		ID:       req.ID,
		UserID:   userID,
		ParentID: req.ParentID,
		Name:     req.Name,
	}

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
	err = mysql.DeferTx(tx, err)
	return category.ID, err
}

func DeleteCategory(userID int64, idStr string) error {
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return err
	}
	userCategoryQuery := query.Category.Where(query.Category.UserID.Eq(userID))
	categoryQuery := userCategoryQuery.Where(query.Category.ID.Eq(id))
	count, err := categoryQuery.Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return &consts.ApiErr{Code: consts.NO_DATA, Msg: "Category not exists."}
	}

	//  在子分类，则不删除
	count, err = query.Category.Where(query.Category.UserID.Eq(userID)).Where(query.Category.ParentID.Eq(id)).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return &consts.ApiErr{Code: consts.DELETE_FAILED, Msg: "Category has subcategories."}
	}

	// 任务列表中存在该分类，不删除
	count, err = query.Task.Where(query.Task.UserID.Eq(userID)).Where(query.Task.CategoryID.Eq(id)).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return &consts.ApiErr{Code: consts.DELETE_FAILED, Msg: "Category is used in tasks."}
	}

	tx := query.Q.Begin()
	info, err := query.Category.Where(query.Category.UserID.Eq(userID)).Where(query.Category.ID.Eq(id)).Delete()
	fmt.Println(info)
	err = mysql.DeferTx(tx, err)
	return err
}
