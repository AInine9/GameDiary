package repository

import "backend/cmd/api/domain/model"

type UserRepository interface {
	FindById(id int) (user *model.User, err error)
	Create(user *model.User) (err error)
	IsNewUser(userId int) bool
}
