package controller

import (
	"github.com/gin-gonic/gin"
	"kyrgyz-bilim/entity"
	"net/http"
)

func UserController(c *gin.Context) {
	param, _ := c.Get("user")
	user := param.(*entity.User)
	c.JSON(http.StatusOK, user)
}
