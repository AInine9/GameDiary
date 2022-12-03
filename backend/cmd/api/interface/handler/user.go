package handler

import (
	"backend/cmd/api/usecase"
)

type UserHandler interface {
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{userUseCase: uu}
}
