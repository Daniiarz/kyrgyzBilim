package repository

import (
	"gorm.io/gorm"
	"kyrgyz-bilim/entity"
	"kyrgyz-bilim/repository/database"
	"time"
)

type UserRepository interface {
	Create(user *entity.User)
	Update(user *entity.User)
	Delete(user *entity.User)
	All() []*entity.User
	GetUserByPhone(phone string) *entity.User
	GetUserById(id int) *entity.User
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		connection: database.DB,
	}
}

func (db *userRepository) GetUserByPhone(phone string) *entity.User {
	user := &entity.User{}
	db.connection.Where("phone_number = ?", phone).First(&user)
	if user.Id == 0 {
		return nil
	}
	return user
}

func (db *userRepository) GetUserById(id int) *entity.User {
	user := &entity.User{}
	db.connection.Where("id = ?", id).First(&user)
	return user
}

func (db *userRepository) All() []*entity.User {
	var users []*entity.User
	db.connection.Find(&users)
	return users
}

func (db *userRepository) Delete(user *entity.User) {
	db.connection.Delete(user)
}

func (db *userRepository) Update(user *entity.User) {
	db.connection.Save(&user)
}

func (db *userRepository) Create(user *entity.User) {
	user.IsActive = true
	user.IsStaff = false
	user.IsSuperuser = false
	user.DateJoined = time.Now()
	user.LastLogin = time.Now()
	db.connection.Save(user)
}
