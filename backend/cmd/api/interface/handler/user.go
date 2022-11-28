package handler

import (
	"backend/cmd/api/usecase"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserHandler interface {
	NewUser(ctx *gin.Context)
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{userUseCase: uu}
}

func (uh userHandler) NewUser(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Query("userId"))
	if err != nil {
		ctx.AbortWithError(500, err)
	}
	userName := ctx.Query("userName")
	err = uh.userUseCase.CreateUser(userId, userName)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}
}
