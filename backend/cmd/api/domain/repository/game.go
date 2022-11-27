package repository

import "backend/cmd/api/domain/model"

type GameRepository interface {
	FindByGameName(gameName string) (game *model.Game, err error)
	Create(game *model.Game) (err error)
	GetGameIdByName(gameName string) (gameId int, err error)
}
