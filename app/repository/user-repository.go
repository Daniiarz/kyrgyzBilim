package repository

import (
	"gorm.io/gorm"
	"kyrgyz-bilim/entity"
	"kyrgyz-bilim/repository/database"
)

type UserRepository interface {
	Create(user *entity.User)
	Update(user *entity.User)
	Delete(user *entity.User)
	All() []*entity.User
	GetUserByPhone(phone string) *entity.User
	GetUserById(id int) *entity.User
}

type repo struct {
	connection *gorm.DB
}

func NewUserRepository() UserRepository {
	return &repo{
		connection: database.DB,
	}
}

func (db *repo) GetUserByPhone(phone string) *entity.User {
	user := &entity.User{}
	database.DB.Where("phone_number = ?", phone).First(&user)
	if user.Id == 0 {
		return nil
	}
	return user
}

func (db *repo) GetUserById(id int) *entity.User {
	user := &entity.User{}
	database.DB.Where("id = ?", id).First(&user)
	return user
}

func (db *repo) All() []*entity.User {
	var users []*entity.User
	db.connection.Find(&users)
	return users
}

func (db *repo) Delete(user *entity.User) {
	db.connection.Delete(user)
}

func (db *repo) Update(user *entity.User) {
	db.connection.Save(&user)
}

func (db *repo) Create(user *entity.User) {
	db.connection.Save(user)
}
