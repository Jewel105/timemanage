package statisticapi

import (
	"gin_study/api"
	"gin_study/api/consts"
	"gin_study/gen/request"
	statisticservice "gin_study/service/statistic_service"

	"github.com/gin-gonic/gin"
)

// @Id GetPieValue
// @Summary 查询分类占比饼图
// @Description 查询分类占比饼图
// @Tags 统计API
// @Accept  json
// @Produce application/json
// @Param token header string false "enjmcvhdwernxhcuvyudfdjfhpq"
// @Param equipment header string false "1234"
// @Param req body request.GetPieValueRequest true "Json"
// @success 200 {object} []response.PieValueResponse "success"
// @Router /statistic/pie [post]
func GetPieValue(c *gin.Context) {
	userID := api.GetUserID(c)
	if userID == 0 {
		return
	}
	req := request.GetPieValueRequest{}
	if !api.ParseJson(c, &req) {
		return
	}
	lang := c.GetString(consts.LANG)

	data, err := statisticservice.GetPieValue(userID, &req, lang)
	api.DealResponse(c, data, err)
}

// @Id GetLineValue
// @Summary 查询分类折线图
// @Description 查询分类折线图
// @Tags 统计API
// @Accept  json
// @Produce application/json
// @Param token header string false "enjmcvhdwernxhcuvyudfdjfhpq"
// @Param equipment header string false "1234"
// @Param req body request.GetLineValueRequest true "Json"
// @success 200 {object} []response.LineValueResponse "success"
// @Router /statistic/line [post]
func GetLineValue(c *gin.Context) {
	userID := api.GetUserID(c)
	if userID == 0 {
		return
	}
	req := request.GetLineValueRequest{}
	if !api.ParseJson(c, &req) {
		return
	}
	data, err := statisticservice.GetLineValue(userID, &req)
	api.DealResponse(c, data, err)
}
