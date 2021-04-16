package routes

import (
	"github.com/gin-gonic/gin"
	"kyrgyz-bilim/controller"
)

func CourseRoutes(rg *gin.RouterGroup) {
	group := rg.Group("/courses")
	{
		group.GET("", controller.Courses)
	}
}
