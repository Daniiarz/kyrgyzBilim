package entity

import (
	"fmt"
	"gorm.io/gorm"
	"os"
)

type Course struct {
	ID          int    `json:"id,omitempty" `
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Section struct {
	ID       int     `json:"id,omitempty"`
	Course   Course  `json:"-" gorm:"foreignKey:CourseID"`
	CourseID int     `json:"-" gorm:"course_id"`
	Name     string  `json:"name,omitempty" gorm:"name"`
	Icon     string  `json:"icon,omitempty" gorm:"icon"`
	Topic    []Topic `json:"topics"`
}

type Topic struct {
	ID             int        `json:"id,omitempty" gorm:"id"`
	SectionID      int        `json:"-"`
	Section        Section    `json:"-" gorm:"foreignKey:SectionID"`
	Name           string     `json:"name,omitempty"`
	TranslatedName string     `json:"translated_name,omitempty" gorm:"translated_name"`
	Icon           string     `json:"icon,omitempty" media:"default"`
	Type           string     `json:"type,omitempty"`
	SubTopic       []SubTopic `json:"sub_topics,omitempty"`
}

type SubTopic struct {
	ID             int    `json:"id,omitempty"`
	TopicId        int    `json:"-	"`
	Topic          Topic  `json:"-" gorm:"foreignKey:TopicId"`
	Text           string `json:"text,omitempty"`
	TranslatedText string `json:"translated_text,omitempty"`
	Audio          string `json:"audio,omitempty"`
	Image          string `json:"image,omitempty"`
}

func (Course) TableName() string {
	return "courses"
}

func (Section) TableName() string {
	return "sections"
}

func (s *Section) AfterFind(tx *gorm.DB) (err error) {
	s.Icon = fmt.Sprintf("%v/%v", os.Getenv("MEDIA_URL"), s.Icon)
	return
}

func (Topic) TableName() string {
	return "topics"
}

func (t *Topic) AfterFind(tx *gorm.DB) (err error) {
	t.Icon = fmt.Sprintf("%v/%v", os.Getenv("MEDIA_URL"), t.Icon)
	return
}

func (SubTopic) TableName() string {
	return "sub_topics"
}

func (s *SubTopic) AfterFind(tx *gorm.DB) (err error) {
	s.Image = fmt.Sprintf("%v/%v", os.Getenv("MEDIA_URL"), s.Image)
	return
}
