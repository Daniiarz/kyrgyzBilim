package service

import (
	"kyrgyz-bilim/entity"
	"kyrgyz-bilim/repository"
)

type CourseService interface {
	AllCourses() []entity.Course
	TopicsById(id int) []entity.Topic
	GetSubtopics(id int, user *entity.User) []entity.SubTopic
	GetSingleTopic(id int) *entity.SubTopic
	CountProgress(user *entity.User, subTopicId int) error
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

func (s courseService) GetSubtopics(id int, user *entity.User) []entity.SubTopic {
	subTopics := s.repository.GetSubtopics(id, user)
	return subTopics
}

func (s courseService) GetSingleTopic(id int) *entity.SubTopic {
	subTopic := s.repository.GetSubtopicById(id)
	return subTopic
}

func (s courseService) CountProgress(user *entity.User, subTopicId int) error {
	subTopic := s.repository.GetSubtopicById(subTopicId)
	err := s.repository.AppendSubTopicToUser(user, subTopic)
	s.repository.RecountUserProgress(user)
	if err != nil {
		return err
	}
	return nil
}
