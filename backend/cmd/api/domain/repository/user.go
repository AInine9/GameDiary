package repository

import "backend/cmd/api/domain/model"

type UserRepository interface {
	Search(name string) ([]*model.User, error)
}
