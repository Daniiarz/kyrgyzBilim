package routes

import (
	"github.com/gin-gonic/gin"
	"kyrgyz-bilim/controller"
)

func UserRoutes(rg *gin.RouterGroup){
	group := rg.Group("users/")
	{
		group.POST("register", controller.Register)
		group.POST("login", controller.Login)
		group.POST("test", controller.TestHashing)
	}
}
