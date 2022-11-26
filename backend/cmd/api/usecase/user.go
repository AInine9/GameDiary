package usecase

import (
	"backend/cmd/api/domain/model"
	"backend/cmd/api/domain/repository"
)

type UserUseCase interface {
	Search(name string) ([]*model.User, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

func (uu userUseCase) Search(name string) (user []*model.User, err error) {
	user, err = uu.userRepository.Search(name)
	if err != nil {
		return nil, err
	}
	return user, nil
}
