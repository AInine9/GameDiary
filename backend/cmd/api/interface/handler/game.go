package handler

import (
	"backend/cmd/api/usecase"
	"github.com/gin-gonic/gin"
)

type GameHandler interface {
	NewGame(ctx *gin.Context)
}

type gameHandler struct {
	gameUseCase usecase.GameUseCase
}

func NewGameHandler(gu usecase.GameUseCase) GameHandler {
	return &gameHandler{gameUseCase: gu}
}

func (gh gameHandler) NewGame(ctx *gin.Context) {
	gameName := ctx.Query("gamename")
	err := gh.gameUseCase.NewGame(gameName)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}
}
