package repository

import (
	"gorm.io/gorm"
	"kyrgyz-bilim/entity"
	"kyrgyz-bilim/repository/database"
)

type CourseRepository interface {
	All() []entity.Course
	GetByID(id int) *entity.Course
	GetSections(id int) []*entity.Section
	GetTopics(id int) []entity.Topic
	GetTopic(id int) *entity.Topic
}

type courseRepository struct {
	connection *gorm.DB
}

func NewCourseRepository() CourseRepository {
	return &courseRepository{
		connection: database.DB,
	}
}

func (db *courseRepository) All() []entity.Course {
	var courses []entity.Course
	db.connection.Find(&courses)
	return courses
}

func (db courseRepository) GetByID(id int) *entity.Course {
	course := &entity.Course{}
	db.connection.Where("id = ?", id).First(&course)
	return course
}

func (db courseRepository) GetSections(id int) []*entity.Section {
	var sections []*entity.Section
	db.connection.Where("course_id = ?", id).Preload("Topic").Find(&sections)
	return sections
}

func (db courseRepository) GetTopics(id int) []entity.Topic {
	var topics []entity.Topic
	db.connection.Where("section_id = ?", id).Find(&topics)
	return topics
}

func (db courseRepository) GetTopic(id int) *entity.Topic {
	topic := &entity.Topic{}
	db.connection.Where("id = ?", id).Preload("SubTopic").Find(&topic)
	return topic
}
