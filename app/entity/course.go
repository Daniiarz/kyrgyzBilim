package entity

type Course struct {
	ID          int    `json:"id,omitempty"`
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
	Icon           string     `json:"icon,omitempty"`
	Type           string     `json:"type,omitempty"`
	SubTopic       []SubTopic `json:"sub_topics"`
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
