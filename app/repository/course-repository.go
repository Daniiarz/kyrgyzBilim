package repository

import (
	"gorm.io/gorm"
	"kyrgyz-bilim/entity"
	"kyrgyz-bilim/repository/database"
	"os"
)

type CourseRepository interface {
	All(user *entity.User) []entity.CourseProgress
	GetByID(id int) *entity.Course
	GetSections(id int) []*entity.Section
	GetTopics(id int) []entity.Topic
	GetTopic(id int) *entity.Topic
	GetSubtopics(id int, user *entity.User) []entity.SubTopic
	AppendSubTopicToUser(user *entity.User, subTopic *entity.SubTopic, courseId int) error
	GetSubtopicById(id int) *entity.SubTopic
}

type courseRepository struct {
	connection *gorm.DB
}

func NewCourseRepository() CourseRepository {
	return &courseRepository{
		connection: database.DB,
	}
}

func (db *courseRepository) All(user *entity.User) []entity.CourseProgress {
	var courses []entity.CourseProgress
	db.connection.Raw(`
		select *,
			   (((select count(1) as count
				  from user_subtopics
				  where user_subtopics.course_id = courses.id and user_id = 1)::float /
				 (select count(1) as count from sub_topics)) * 100)::int as progress
		from courses
	`, user.Id).Scan(&courses)
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
		"SELECT s.id, s.text, s.translated_text, concat(?::text, '/'::text, s.audio) as audio, s.image, s.order, CAST(u.id::int as BOOLEAN) AS completed "+
		"FROM sub_topics as s "+
		"LEFT  JOIN user_subtopics as us "+
		"ON s.id=us.sub_topic_id "+
		"AND us.user_id = ? "+
		"LEFT JOIN users as u "+
		"ON us.user_id=u.id "+
		"WHERE s.topic_id = ?", os.Getenv("MEDIA_URL"), user.Id, id).Scan(&subTopics)
	return subTopics
}

func (db courseRepository) GetSubtopicById(id int) *entity.SubTopic {
	subTopic := &entity.SubTopic{}
	db.connection.Where("id = ?", id).Find(&subTopic)
	return subTopic
}

func (db courseRepository) AppendSubTopicToUser(user *entity.User, subTopic *entity.SubTopic, courseId int) error {
	userSubtopic := entity.UserSubtopic{
		UserId:     user.Id,
		SubTopicId: subTopic.ID,
		CourseId:   courseId,
	}
	err := db.connection.Create(&userSubtopic)
	if err != nil {
	}
	return nil
}
