package model

import (
	"time"
)

type Diary struct {
	Id             int       `json:"id" db:"id"`
	UserId         int       `json:"user_id" db:"user_id"`
	GameId         int       `json:"game_id" db:"game_id"`
	StartPlayingAt time.Time `json:"start_playing_at" db:"start_playing_at"`
	EndPlayingAt   time.Time `json:"end_playing_at" db:"end_playing_at"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}
