package routes

import (
	"github.com/gin-gonic/gin"
	"kyrgyz-bilim/controller"
)

func CourseRoutes(rg *gin.RouterGroup) {
	group := rg.Group("/courses")
	{
		group.GET("", controller.ListCourses)
		group.GET("/:courseId/sections", controller.ListSections)
		group.GET("/:courseId/sections/:sectionId/topics", controller.ListTopics)
	}
}
