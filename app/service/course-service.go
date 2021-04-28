package service

import (
	"kyrgyz-bilim/entity"
	"kyrgyz-bilim/repository"
)

type CourseService interface {
	AllCourses() []entity.Course
	TopicsById(id int) []entity.Topic
}

type courseService struct {
	repository repository.CourseRepository
}

func NewCourseService() CourseService {
	repo := repository.NewCourseRepository()
	return courseService{
		repository: repo,
	}
}

func (s courseService) AllCourses() []entity.Course {
	return s.repository.All()
}

func (s courseService) TopicsById(id int) []entity.Topic {
	topics := s.repository.GetTopics(id)
	return topics
}
