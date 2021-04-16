package entity

type Course struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Section struct {
	ID       int    `json:"id,omitempty"`
	Course   Course `json:"course" gorm:"foreignKey:CourseID"`
	CourseID int    `json:"course_id,omitempty" gorm:"course_id"`
	Name     string `json:"name,omitempty" gorm:"name"`
	Icon     string `json:"icon,omitempty" gorm:"icon"`
}

type Topic struct {
	ID             int     `json:"id,omitempty" gorm:"id"`
	Section        Section `json:"section" gorm:"foreignKey:SectionID"`
	SectionID      int     `json:"section_id,omitempty" gorm:"section_id"`
	Name           string  `json:"name,omitempty" gorm:"name"`
	TranslatedName string  `json:"translated_name,omitempty" gorm:"translated_name"`
	Icon           string  `json:"icon,omitempty" gorm:"icon"`
	Type           string  `json:"type,omitempty" gorm:"type"`
}

type SubTopic struct {
	ID             int    `json:"id,omitempty"`
	TopicId        int    `json:"topic_id,omitempty"`
	Topic          Topic  `json:"topic" gorm:"foreignKey:TopicID"`
	Text           string `json:"text,omitempty"`
	TranslatedText string `json:"translated_text,omitempty"`
	Audio          string `json:"audio,omitempty"`
	Image          string `json:"image,omitempty"`
}

func (Course) TableName() string {
	return "courses"
}
