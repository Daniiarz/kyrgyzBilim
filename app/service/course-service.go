package service

import (
	"kyrgyz-bilim/entity"
	"kyrgyz-bilim/repository"
)

type CourseService interface {
	TopicsById(id int) []entity.Topic
	GetSubtopics(id int, user *entity.User) []entity.SubTopic
	GetSingleTopic(id int) *entity.SubTopic
	CountProgress(user *entity.User, subTopicId int, courseId int) error
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

func (s courseService) TopicsById(id int) []entity.Topic {
	topics := s.repository.GetTopics(id)
	return topics
}

func (s courseService) GetSubtopics(id int, user *entity.User) []entity.SubTopic {
	subTopics := s.repository.GetSubtopics(id, user)
	return subTopics
}

func (s courseService) GetSingleTopic(id int) *entity.SubTopic {
	subTopic := s.repository.GetSubtopicById(id)
	return subTopic
}

func (s courseService) CountProgress(user *entity.User, subTopicId int, courseId int) error {
	subTopic := s.repository.GetSubtopicById(subTopicId)
	err := s.repository.AppendSubTopicToUser(user, subTopic, courseId)
	if err != nil {
		println(err.Error())
	}
	return nil
}
