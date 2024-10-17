package categoryapi

import (
	"gin_study/api"
	"gin_study/gen/request"
	categoryservice "gin_study/service/category_service"

	"github.com/gin-gonic/gin"
)

// @Id CategoryGetList
// @Summary 查询分类列表
// @Description 查询分类列表
// @Tags 分类API
// @Accept  json
// @Produce application/json
// @Param token header string false "enjmcvhdwernxhcuvyudfdjfhpq"
// @Param ParentID query int64 false "11"
// @success 200 {object} []response.CategoriesRespose "success"
// @Router /categories/list [get]
func GetList(c *gin.Context) {
	userID := api.GetUserID(c)
	if userID == 0 {
		return
	}
	req := request.GetCategoriesRequest{}
	if !api.ParseQuery(c, &req) {
		return
	}
	categories, err := categoryservice.GetList(userID, &req)
	api.DealResponse(c, categories, err)
}

// @Id SaveCategory
// @Summary 保存或修改分类
// @Description 保存或修改分类
// @Tags 分类API
// @Accept  json
// @Produce application/json
// @Param token header string false "enjmcvhdwernxhcuvyudfdjfhpq"
// @Param req body request.SaveCategoryRequest true "Json"
// @success 200 int64 categoryID "success"
// @Router  /categories/save [post]
func SaveCategory(c *gin.Context) {
	userID := api.GetUserID(c)
	if userID == 0 {
		return
	}
	req := request.SaveCategoryRequest{}
	if !api.ParseJson(c, &req) {
		return
	}
	categoryID, err := categoryservice.SaveCategory(userID, &req)
	api.DealResponse(c, categoryID, err)
}

// @Id DeleteCategory
// @Summary 删除分类
// @Description 删除分类
// @Tags 分类API
// @Accept  json
// @Produce application/json
// @Param token header string false "enjmcvhdwernxhcuvyudfdjfhpq"
// @Param id path int true "分类ID"
// @success 200 boolean ture "success"
// @Router  /categories/delete/:id [post]
func DeleteCategory(c *gin.Context) {
	userID := api.GetUserID(c)
	if userID == 0 {
		return
	}
	idStr := c.Param("id")
	err := categoryservice.DeleteCategory(userID, idStr)
	api.DealResponse(c, true, err)
}
