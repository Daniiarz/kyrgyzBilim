package routes

import (
	"github.com/gin-gonic/gin"
	"kyrgyz-bilim/controller"
	"kyrgyz-bilim/utils"
)

func UserRouters(rg *gin.RouterGroup) {
	group := rg.Group("/user")
	{
		group.GET("", utils.AuthMiddleware(), controller.UserController)
	}
}
