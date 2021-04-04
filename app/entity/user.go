package entity

type User struct {
	Id          uint64 `json:"id" gorm:"primaryKey;autoIncrement" binding:"-"`
	FirstName   string `json:"first_name" gorm:"type:varchar(20)" binding:"required,min=3,max=20"`
	LastName    string `json:"last_name" gorm:"type:varchar(20)" binding:"required,min=3,max=20"`
	PhoneNumber string `json:"phone_number" gorm:"uniqueIndex" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

func (User) TableName() string {
	return "users"
}
