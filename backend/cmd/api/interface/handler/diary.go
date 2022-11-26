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
}

func NewDiaryHandler(du usecase.DiaryUseCase) DiaryHandler {
	return &diaryHandler{diaryUseCase: du}
}

func (dh diaryHandler) Index(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (dh diaryHandler) StartPlaying(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Query("userId"))
	gameId, _ := strconv.Atoi(ctx.Query("gameId"))
	err := dh.diaryUseCase.Create(userId, gameId)
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
