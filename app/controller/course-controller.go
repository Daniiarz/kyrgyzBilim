package controller

import (
	"github.com/gin-gonic/gin"
	"kyrgyz-bilim/repository"
	"net/http"
	"strconv"
)

func ListCourses(c *gin.Context) {
	connection := repository.NewCourseRepository()
	courses := connection.All()
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
	connection := repository.NewCourseRepository()
	id := c.Param("id")
	if sectionId, err := strconv.Atoi(id); err == nil {
		topics := connection.GetTopics(sectionId)
		c.JSON(http.StatusOK, topics)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be integer"})
	}
}

func DetailedTopic(c *gin.Context) {
	connection := repository.NewCourseRepository()
	id := c.Param("id")
	if sectionId, err := strconv.Atoi(id); err == nil {
		topic := connection.GetTopic(sectionId)
		c.JSON(http.StatusOK, topic)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be integer"})
	}
}
