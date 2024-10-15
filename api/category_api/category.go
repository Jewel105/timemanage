package categoryapi

import (
	"gin_study/api"
	"gin_study/gen/request"
	categoryservice "gin_study/service/category_service"

	"github.com/gin-gonic/gin"
)

// @Id GetList
// @Summary 查询brc20铭刻详情
// @Description 查询brc20铭刻详情
// @Tags BRC20API
// @Accept  json
// @Produce application/json
// @Param token header string false "869099058070000"
// @Param parentID query int64 false "1"
// @success 200 {object} []models.Category "success"
// @Router /brc20/trade/detail [get]
// @Security Bearer
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

func DeleteCategory(c *gin.Context) {
	userID := api.GetUserID(c)
	if userID == 0 {
		return
	}
	idStr := c.Param("id")
	err := categoryservice.DeleteCategory(userID, idStr)
	api.DealResponse(c, true, err)
}
