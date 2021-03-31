package entity

import (
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        uint64 `json:"id" gorm:"primaryKey;autoIncrement" binding:"-"`
	FirstName string `json:"first_name" gorm:"type:varchar(20)" binding:"required"`
	LastName  string `json:"last_name" gorm:"type:varchar(20)" binding:"required"`
	Email     string `json:"email" gorm:"uniqueIndex" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

func (User) TableName() string {
	return "users"
}

func (user *User) SetPassword() {
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	base64EncodedPasswordHash := base64.URLEncoding.EncodeToString(hashedPass)
	user.Password = base64EncodedPasswordHash
}