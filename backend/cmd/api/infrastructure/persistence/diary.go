package persistence

import (
	"backend/cmd/api/domain/model"
	"backend/cmd/api/domain/repository"
	"github.com/jmoiron/sqlx"
	"time"
)

type diaryPersistence struct {
	Conn *sqlx.DB
}

func NewDiaryPersistence(conn *sqlx.DB) repository.DiaryRepository {
	return &diaryPersistence{Conn: conn}
}

func (dp *diaryPersistence) FindAll() (diaries []*model.Diary, err error) {
	db := dp.Conn

	if err := db.Select(&diaries, "SELECT * FROM diaries"); err != nil {
		return nil, err
	}

	return diaries, nil
}

func (dp *diaryPersistence) FindWritingDiary(userId int) (diary *model.Diary, err error) {
	db := dp.Conn

	diary = new(model.Diary)
	err = db.QueryRowx("SELECT * FROM diaries WHERE user_id = ? AND start_playing_at = end_playing_at LIMIT 1", userId).StructScan(diary)
	if err != nil {
		return nil, err
	}
	return diary, nil
}

func (dp *diaryPersistence) Create(diary *model.Diary) (err error) {
	db := dp.Conn
	now := time.Now()

	_, err = db.Exec(
		"INSERT INTO diaries (id, user_id, game_id, start_playing_at, end_playing_at, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)",
		diary.Id, diary.UserId, diary.GameId, diary.StartPlayingAt, diary.EndPlayingAt, now, now)
	if err != nil {
		return err
	}

	return nil
}

func (dp *diaryPersistence) Update(diary *model.Diary) (err error) {
	db := dp.Conn
	now := time.Now()

	_, err = db.Exec("UPDATE diaries SET user_id = ?, game_id = ?, start_playing_at = ?, end_playing_at = ?, created_at = ?, updated_at = ? WHERE id = ?",
		diary.UserId, diary.GameId, diary.StartPlayingAt, diary.EndPlayingAt, diary.CreatedAt, now, diary.Id)
	if err != nil {
		return err
	}

	return nil
}
