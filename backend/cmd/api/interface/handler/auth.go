package handler

import (
	"backend/cmd/api/usecase"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"github.com/stretchr/objx"
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
	}
	authCookieValue := objx.New(map[string]interface{}{
		"user_id":   userId,
		"user_name": userName,
	}).MustBase64()
	ctx.SetCookie("auth", authCookieValue, 86400*30, "/", "localhost", false, true)
}

func (ah authHandler) BeginAuth(ctx *gin.Context) {
	gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
}

func (ah authHandler) Logout(ctx *gin.Context) {
	gothic.Logout(ctx.Writer, ctx.Request)
	ctx.SetCookie("auth", "", -1, "/", "localhost", false, true)
}
