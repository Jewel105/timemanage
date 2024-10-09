package api

import "github.com/gin-gonic/gin"

type TaskController struct{}

func (t TaskController) GetList(c *gin.Context) {
	if _, exists := c.Get(USER_ID); !exists {
		ReturnResponse(c, TOKEN_INVALID, "user not found")
		return
	}
}
