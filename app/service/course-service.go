package service

import (
	"kyrgyz-bilim/entity"
	"kyrgyz-bilim/repository"
)

type CourseService interface {
	AllCourses() []*entity.Course
}

type courseService struct {
	repository repository.CourseRepository
}

func (s courseService) AllCourses() []*entity.Course {
	return s.repository.All()
}
