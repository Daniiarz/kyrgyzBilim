package entity

type User struct {
	Id             uint64 `json:"id" gorm:"primaryKey;autoIncrement" form:"-" binding:"-"`
	FirstName      string `json:"first_name" gorm:"type:varchar(20)" form:"first_name" binding:"required,min=3,max=20"`
	LastName       string `json:"last_name" gorm:"type:varchar(20)" form:"last_name" binding:"required,min=3,max=20"`
	PhoneNumber    string `json:"phone_number" gorm:"uniqueIndex" form:"phone_number" binding:"required"`
	Password       string `json:"password" form:"password" binding:"required"`
	ProfilePicture string `json:"profile_picture" gorm:"type:text" binding:"-"`
}

func (User) TableName() string {
	return "users"
}
