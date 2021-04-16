package service

import (
	"kyrgyz-bilim/entity"
	"kyrgyz-bilim/repository"
)

var repo = repository.NewCourseRepository()

type CourseService interface {
	AllCourses() []*entity.Course
}

type courseService struct {
	repository repository.CourseRepository
}

func NewCourseService() CourseService {
	return &courseService{
		repository: repo,
	}
}

func (s courseService) AllCourses() []*entity.Course {
	return s.repository.All()
}
