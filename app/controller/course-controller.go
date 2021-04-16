package controller

import (
	"github.com/gin-gonic/gin"
	"kyrgyz-bilim/service"
	"net/http"
)

var courseService = service.NewCourseService()

func Courses(c *gin.Context) {
	courses := courseService.AllCourses()
	c.JSON(http.StatusOK, courses)
}
