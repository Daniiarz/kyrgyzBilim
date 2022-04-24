package entity

import (
	"fmt"
	"gorm.io/gorm"
	"os"
)

type Course struct {
	ID          int    `json:"id," `
	Name        string `json:"name,"`
	Description string `json:"description,"`
	Progress    int    `json:"progress" gorm:"-"`
}

type CourseProgress struct {
	ID          int    `json:"id," `
	Name        string `json:"name,"`
	Description string `json:"description,"`
	Progress    int    `json:"progress"`
}

type Section struct {
	ID       int     `json:"id,omitempty"`
	Course   Course  `json:"-" gorm:"foreignKey:CourseID"`
	CourseID int     `json:"-" gorm:"course_id"`
	Name     string  `json:"name," gorm:"name"`
	Icon     string  `json:"icon," gorm:"icon"`
	Type     string  `json:"type,"`
	Topic    []Topic `json:"topics"`
}

type Topic struct {
	ID             int        `json:"id,omitempty" gorm:"id"`
	SectionID      int        `json:"-"`
	Section        Section    `json:"-" gorm:"foreignKey:SectionID"`
	Name           string     `json:"name,"`
	TranslatedName string     `json:"translated_name," gorm:"translated_name"`
	Icon           string     `json:"icon," media:"default"`
	Type           string     `json:"type,"`
	SubTopic       []SubTopic `json:"sub_topics,"`
}

type SubTopic struct {
	ID             int    `json:"id,omitempty"`
	TopicId        int    `json:"-"`
	Topic          Topic  `json:"-" gorm:"foreignKey:TopicId"`
	Text           string `json:"text,"`
	TranslatedText string `json:"translated_text,"`
	Audio          string `json:"audio,"`
	Image          string `json:"image,"`
	Order          int    `json:"order"`
	Completed      bool   `json:"completed"`
}

func (Course) TableName() string {
	return "courses"
}

func (Section) TableName() string {
	return "sections"
}

func (s *Section) AfterFind(tx *gorm.DB) (err error) {
	mediaUrl := os.Getenv("MEDIA_URL")
	if s.Icon != "" {
		s.Icon = fmt.Sprintf("%v/%v", mediaUrl, s.Icon)
	}
	return
}

func (Topic) TableName() string {
	return "topics"
}

func (t *Topic) AfterFind(tx *gorm.DB) (err error) {
	mediaUrl := os.Getenv("MEDIA_URL")
	if t.Icon != "" {
		t.Icon = fmt.Sprintf("%v/%v", mediaUrl, t.Icon)
	}
	return
}

func (SubTopic) TableName() string {
	return "sub_topics"
}

func (s *SubTopic) AfterFind(tx *gorm.DB) (err error) {
	mediaUrl := os.Getenv("MEDIA_URL")
	if s.Image != "" {
		s.Image = fmt.Sprintf("%v/%v", mediaUrl, s.Image)
	}
	if s.Audio != "" {
		s.Audio = fmt.Sprintf("%v/%v", mediaUrl, s.Audio)
	}
	return
}
