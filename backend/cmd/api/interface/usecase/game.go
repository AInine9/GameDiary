package usecase

import (
	"backend/cmd/api/domain/model"
	"backend/cmd/api/domain/repository"
)

type GameUseCase interface {
	NewGame(gameName string) (err error)
	GetGameIdByName(gameName string) (gameId int, err error)
}

type gameUseCase struct {
	gameRepository repository.GameRepository
}

func NewGameUseCase(gr repository.GameRepository) GameUseCase {
	return &gameUseCase{gameRepository: gr}
}

func (gu gameUseCase) NewGame(gameName string) (err error) {
	game := &model.Game{
		Name: gameName,
	}

	err = gu.gameRepository.Create(game)
	if err != nil {
		return err
	}
	return nil
}

func (gu gameUseCase) GetGameIdByName(gameName string) (gameId int, err error) {
	gameId, err = gu.gameRepository.GetGameIdByName(gameName)
	if err != nil {
		return 0, err
	}
	return gameId, nil
}
