package api

import "github.com/gin-gonic/gin"

type OrderController struct{}

func (o OrderController) GetInfo(c *gin.Context) {
	ReturnResponse(c, SUCCESS, true)
}

func (o OrderController) GetList(c *gin.Context) {
	ReturnResponse(c, SUCCESS, true)

}
