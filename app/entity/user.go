package entity

import (
	"fmt"
	"gorm.io/gorm"
	"os"
	"time"
)

type User struct {
	Id             uint64    `json:"id" gorm:"primaryKey;autoIncrement" form:"-" binding:"-"`
	FirstName      string    `json:"first_name" gorm:"type:varchar(20)" form:"first_name" binding:"required,min=3,max=20"`
	LastName       string    `json:"last_name" gorm:"type:varchar(20)" form:"last_name" binding:"required,min=3,max=20"`
	PhoneNumber    string    `json:"phone_number" gorm:"type:varchar(30);uniqueIndex" form:"phone_number" binding:"required"`
	Password       string    `json:"-" form:"password" binding:"required"`
	ProfilePicture string    `json:"profile_picture" gorm:"type:varchar(255)" binding:"-"`
	Level          uint      `json:"level" binding:"-"`
	LastLogin      time.Time `json:"-" binding:"-"`
	DateJoined     time.Time `json:"-" binding:"-"`
	IsActive       bool      `json:"-" binding:"-"`
	IsStaff        bool      `json:"-" binding:"-"`
	IsSuperuser    bool      `json:"-" binding:"-"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	u.ProfilePicture = fmt.Sprintf("%v/%v", os.Getenv("MEDIA_URL"), u.ProfilePicture)
	return
}
