package repository

import (
	"gorm.io/gorm"
	"kyrgyz-bilim/entity"
	"kyrgyz-bilim/repository/database"
)

type CourseRepository interface {
	All() []*entity.Course
}

type courseRepository struct {
	connection *gorm.DB
}

func NewCourseRepository() CourseRepository {
	return &courseRepository{
		connection: database.DB,
	}
}

func (db *courseRepository) All() []*entity.Course {
	var courses []*entity.Course
	db.connection.Find(&courses)
	return courses
}
