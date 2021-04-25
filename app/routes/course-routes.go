package routes

import (
	"github.com/gin-gonic/gin"
	"kyrgyz-bilim/controller"
)

func CourseRoutes(rg *gin.RouterGroup) {
	group := rg.Group("")
	{
		group.GET("courses", controller.ListCourses)
		group.GET("courses/:id/sections", controller.ListSections)
		group.GET("sections/:id/topics", controller.ListTopics)
		group.GET("topics/:id", controller.DetailedTopic)
	}
}
