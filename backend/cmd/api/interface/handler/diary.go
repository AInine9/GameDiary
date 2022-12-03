package handler

import (
	"backend/cmd/api/usecase"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type DiaryHandler interface {
	Index(*gin.Context)
	StartPlaying(*gin.Context)
	EndPlaying(*gin.Context)
}

type diaryHandler struct {
	diaryUseCase usecase.DiaryUseCase
	gameUseCase  usecase.GameUseCase
	userUseCase  usecase.UserUseCase
}

func NewDiaryHandler(du usecase.DiaryUseCase, gu usecase.GameUseCase, uu usecase.UserUseCase) DiaryHandler {
	return &diaryHandler{diaryUseCase: du, gameUseCase: gu, userUseCase: uu}
}

func (dh diaryHandler) Index(ctx *gin.Context) {
}

func (dh diaryHandler) StartPlaying(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Query("userId"))
	if dh.userUseCase.CheckIsNewUser(userId) {
		ctx.AbortWithError(401, errors.New(fmt.Sprintf("user: %d is not registered.", userId)))
		return
	}

	gameName := ctx.Query("gameName")
	gameId, err := dh.gameUseCase.GetGameIdByName(gameName)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}
	if gameId == 0 {
		err = dh.gameUseCase.NewGame(gameName)
		if err != nil {
			ctx.AbortWithError(500, err)
		}
		gameId, err = dh.gameUseCase.GetGameIdByName(gameName)
		if err != nil {
			ctx.AbortWithError(500, err)
		}
	}
	err = dh.diaryUseCase.Create(userId, gameId)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}
}

func (dh diaryHandler) EndPlaying(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Query("userId"))
	if dh.userUseCase.CheckIsNewUser(userId) {
		ctx.AbortWithError(401, errors.New(fmt.Sprintf("user: %d is not registered.", userId)))
		return
	}

	err := dh.diaryUseCase.EndPlaying(userId)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}
}
