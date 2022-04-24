package controller

import (
	"github.com/gin-gonic/gin"
	"kyrgyz-bilim/entity"
	"kyrgyz-bilim/repository"
	"kyrgyz-bilim/service"
	"net/http"
	"strconv"
)

func ListCourses(c *gin.Context) {
	connection := repository.NewCourseRepository()
	param, _ := c.Get("user")
	user := param.(*entity.User)
	courses := connection.All(user)
	c.JSON(http.StatusOK, courses)
}

func ListSections(c *gin.Context) {
	connection := repository.NewCourseRepository()
	id := c.Param("id")
	if courseId, err := strconv.Atoi(id); err == nil {
		sections := connection.GetSections(courseId)
		c.JSON(http.StatusOK, sections)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be integer"})
	}
}

func ListTopics(c *gin.Context) {
	id := c.Param("id")
	if sectionId, err := strconv.Atoi(id); err == nil {
		newService := service.NewCourseService()
		topics := newService.TopicsById(sectionId)
		c.JSON(http.StatusOK, topics)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be integer"})
	}
}

func DetailedTopic(c *gin.Context) {
	id := c.Param("id")
	if sectionId, err := strconv.Atoi(id); err == nil {
		newService := service.NewCourseService()
		param, _ := c.Get("user")
		user := param.(*entity.User)
		subTopics := newService.GetSubtopics(sectionId, user)
		c.JSON(http.StatusOK, subTopics)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be integer"})
	}
}

func CountProgress(c *gin.Context) {
	id := c.Param("id")
	if subTopicId, err := strconv.Atoi(id); err == nil {
		param, _ := c.Get("user")
		user := param.(*entity.User)
		courseId := &entity.CourseId{}
		obj, ok := service.DataBind(c, courseId)
		if !ok {
			c.JSON(http.StatusBadRequest, obj.(gin.H))
			return
		}
		newService := service.NewCourseService()
		err = newService.CountProgress(user, subTopicId, obj.(*entity.CourseId).CourseId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.Error{Err: err})
		}
		c.JSON(http.StatusOK, gin.H{})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be integer"})
	}
}
