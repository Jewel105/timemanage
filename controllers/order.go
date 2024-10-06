package controllers

import "github.com/gin-gonic/gin"

type OrderController struct{}

func (o OrderController) GetInfo(c *gin.Context) {
	ReturnSuccess(c, "0", "success", "value", 3)
}

func (o OrderController) GetList(c *gin.Context) {
	ReturnError(c, "4001", "fail")
}
