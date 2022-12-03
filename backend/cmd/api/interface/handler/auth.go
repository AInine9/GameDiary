package handler

import (
	"backend/cmd/api/usecase"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"strconv"
)

type AuthHandler interface {
	CompleteAuth(ctx *gin.Context)
	BeginAuth(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type authHandler struct {
	userUseCase usecase.UserUseCase
}

func NewAuthHandler(uu usecase.UserUseCase) AuthHandler {
	return &authHandler{userUseCase: uu}
}

func (ah authHandler) CompleteAuth(ctx *gin.Context) {
	user, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}
	userId, _ := strconv.Atoi(user.UserID)
	userName := user.Name
	if ah.userUseCase.CheckIsNewUser(userId) {
		err = ah.userUseCase.CreateUser(userId, userName)
		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}
		// TODO 新規ユーザ登録後の処理
	} else {
		// TODO 既存ユーザログイン後の処理
	}
}

func (ah authHandler) BeginAuth(ctx *gin.Context) {
	gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
}

func (ah authHandler) Logout(ctx *gin.Context) {
	gothic.Logout(ctx.Writer, ctx.Request)
	// TODO ログアウト後の処理
}
