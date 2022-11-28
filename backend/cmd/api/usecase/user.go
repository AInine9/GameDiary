package usecase

import (
	"backend/cmd/api/domain/model"
	"backend/cmd/api/domain/repository"
)

type UserUseCase interface {
	CreateUser(userId int, userName string) (err error)
	CheckIsNewUser(userId int) bool
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{userRepository: ur}
}

func (uu userUseCase) CreateUser(userId int, userName string) (err error) {
	user := &model.User{
		Id:       userId,
		UserName: userName,
	}
	err = uu.userRepository.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (uu userUseCase) CheckIsNewUser(userId int) bool {
	return uu.userRepository.IsNewUser(userId)
}
