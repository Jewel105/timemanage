package categoryservice

import (
	"fmt"
	"gin_study/api/consts"
	"gin_study/gen/models"
	"gin_study/gen/mysql"
	"gin_study/gen/query"
	"gin_study/gen/request"
	"gin_study/gen/response"
	"gin_study/language"
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

func DeleteCategory(userID int64, idStr, lang string) error {
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return err
	}
	userCategoryQuery := query.Category.Where(query.Category.UserID.Eq(userID))
	categoryQuery := userCategoryQuery.Where(query.Category.ID.Eq(id))
	curCategory, err := categoryQuery.First()
	if err != nil {
		return &consts.ApiErr{Code: consts.NO_DATA, Msg: language.GetLocale(lang, "NoCategory")}
	}

	//  在子分类，则不删除
	count, err := query.Category.Where(query.Category.UserID.Eq(userID)).Where(query.Category.ParentID.Eq(id)).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return &consts.ApiErr{Code: consts.DELETE_FAILED, Msg: language.GetLocale(lang, "CategoryHasSub")}
	}

	tx := query.Q.Begin()

	if curCategory.ParentID != 0 {
		// 如果不是顶级分类，将任务列列表中的path和categoryID改为父分类
		_, err = query.Task.Where(query.Task.UserID.Eq(userID)).Where(query.Task.CategoryID.Eq(id)).Updates(models.Task{CategoryPath: curCategory.Path, CategoryID: curCategory.ParentID})
		if err != nil {
			return err
		}
	} else {
		// 如果是顶级分类，有任务存在则不能删除
		count, err = query.Task.Where(query.Task.UserID.Eq(userID)).Where(query.Task.CategoryID.Eq(id)).Count()
		if err != nil {
			return err
		}
		if count > 0 {
			return &consts.ApiErr{Code: consts.DELETE_FAILED, Msg: language.GetLocale(lang, "CategoryUsed")}
		}
	}

	// 删除分类
	info, err := query.Category.Where(query.Category.UserID.Eq(userID)).Where(query.Category.ID.Eq(id)).Delete()
	fmt.Println(info)
	err = mysql.DeferTx(tx, err)
	return err
}
