package handler

import (
	"backend/cmd/api/interface/usecase"
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
}

func NewDiaryHandler(du usecase.DiaryUseCase, gu usecase.GameUseCase) DiaryHandler {
	return &diaryHandler{diaryUseCase: du, gameUseCase: gu}
}

func (dh diaryHandler) Index(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (dh diaryHandler) StartPlaying(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Query("userId"))
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
	}
	err = dh.diaryUseCase.Create(userId, gameId)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}
}

func (dh diaryHandler) EndPlaying(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Query("userId"))
	err := dh.diaryUseCase.EndPlaying(userId)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}
}
