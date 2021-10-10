package routes

import (
	"github.com/gin-gonic/gin"
	"kyrgyz-bilim/controller"
	"kyrgyz-bilim/utils"
)

func CourseRoutes(rg *gin.RouterGroup) {
	group := rg.Group("")
	{
		group.GET("courses", controller.ListCourses)
		group.GET("courses/:id/sections", controller.ListSections)
		group.GET("sections/:id/topics", controller.ListTopics)
		group.GET("topics/:id", utils.AuthMiddleware(), controller.DetailedTopic)
		group.POST("subtopics/:id/count-progress", utils.AuthMiddleware(), controller.CountProgress)
	}
}
