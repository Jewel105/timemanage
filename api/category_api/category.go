package categoryapi

import (
	"gin_study/api"
	"gin_study/gen/request"
	categoryservice "gin_study/service/category_service"

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
