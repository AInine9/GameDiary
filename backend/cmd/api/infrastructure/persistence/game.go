package persistence

import (
	"backend/cmd/api/domain/model"
	"backend/cmd/api/domain/repository"
	"github.com/jmoiron/sqlx"
	"time"
)

type gamePersistence struct {
	Conn *sqlx.DB
}

func NewGamePersistence(conn *sqlx.DB) repository.GameRepository {
	return &gamePersistence{Conn: conn}
}

func (gp *gamePersistence) FindByGameName(gameName string) (game *model.Game, err error) {
	db := gp.Conn

	if err = db.Select(&game, "SELECT * FROM games WHERE name = ? LIMIT 1", gameName); err != nil {
		return nil, err
	}

	return game, nil
}

func (gp *gamePersistence) Create(game *model.Game) (err error) {
	db := gp.Conn
	now := time.Now()

	_, err = db.Exec("INSERT INTO games (name, created_at, updated_at) VALUES (?, ?, ?)",
		game.Name, now, now)
	if err != nil {
		return err
	}
	return nil
}

func (gp *gamePersistence) GetGameIdByName(gameName string) (gameId int, err error) {
	db := gp.Conn

	row := db.QueryRowx("SELECT id FROM games WHERE name = ?", gameName)
	err = row.Scan(&gameId)
	if err != nil {
		return 0, nil
	}
	return gameId, nil
}
