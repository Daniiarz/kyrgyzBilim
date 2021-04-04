package entity

type User struct {
	Id          uint64 `json:"id" gorm:"primaryKey;autoIncrement" binding:"-"`
	FirstName   string `json:"first_name" gorm:"type:varchar(20)" binding:"required"`
	LastName    string `json:"last_name" gorm:"type:varchar(20)" binding:"required"`
	PhoneNumber string `json:"phone_number" gorm:"uniqueIndex" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

func (User) TableName() string {
	return "users"
}
