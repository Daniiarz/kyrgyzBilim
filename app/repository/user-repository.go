package repository

import (
	"kyrgyz-bilim/entity"
)

type UserRepository interface {
	Save(user entity.User)
	Update(user entity.User)
	Delete(user entity.User)
	All(user entity.User)
}

//func NewUserRepository() UserRepository {
//
//}
