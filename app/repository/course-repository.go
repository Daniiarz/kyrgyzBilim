package repository

import (
	"fmt"
	"gorm.io/gorm"
	"kyrgyz-bilim/entity"
	"kyrgyz-bilim/repository/database"
	"log"
)

type CourseRepository interface {
	All() []entity.Course
	GetByID(id int) *entity.Course
	GetSections(id int) []*entity.Section
	GetTopics(id int) []entity.Topic
	GetTopic(id int) *entity.Topic
	GetSubtopics(id int, user *entity.User) []entity.SubTopic
	AppendSubTopicToUser(user *entity.User, subTopic *entity.SubTopic) error
	GetSubtopicById(id int) *entity.SubTopic
	RecountUserProgress(user *entity.User)
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

func (db courseRepository) GetSubtopics(id int, user *entity.User) []entity.SubTopic {
	var subTopics []entity.SubTopic
	db.connection.Raw(""+
		"SELECT s.id, s.text, s.translated_text, s.audio, s.image, s.order, CAST(u.id::int as BOOLEAN) AS completed "+
		"FROM sub_topics as s "+
		"LEFT  JOIN user_subtopics as us "+
		"ON s.id=us.sub_topic_id "+
		"AND us.user_id = ? "+
		"LEFT JOIN users as u "+
		"ON us.user_id=u.id "+
		"WHERE s.id = ?", user.Id, id).Scan(&subTopics)
	return subTopics
}

func (db courseRepository) GetSubtopicById(id int) *entity.SubTopic {
	subTopic := &entity.SubTopic{}
	db.connection.Where("id = ?", id).Find(&subTopic)
	return subTopic
}

func (db courseRepository) AppendSubTopicToUser(user *entity.User, subTopic *entity.SubTopic) error {
	err := db.connection.Model(&user).Association("SubTopics").Append(subTopic)
	if err != nil {
		log.Fatal(err.Error())
	}
	return nil
}

func (db courseRepository) RecountUserProgress(user *entity.User) {
	var subTopicCount int64
	userSubTopics := db.connection.Model(&user).Association("SubTopics").Count()
	tableName := entity.SubTopic{}.TableName()
	db.connection.Table(tableName).Count(&subTopicCount)
	fmt.Println(subTopicCount)
	fmt.Println(userSubTopics)
	user.Progress = float64(userSubTopics) / float64(subTopicCount) * 100
	db.connection.Save(&user)
}
