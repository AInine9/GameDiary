package persistence

import (
	"backend/cmd/api/domain/model"
	"backend/cmd/api/domain/repository"
	"github.com/jmoiron/sqlx"
	"time"
)

type userPersistence struct {
	Conn *sqlx.DB
}

func NewUserPersistence(conn *sqlx.DB) repository.UserRepository {
	return &userPersistence{Conn: conn}
}

func (up *userPersistence) FindById(id int) (user *model.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (up *userPersistence) Create(user *model.User) (err error) {
	db := up.Conn
	now := time.Now()

	_, err = db.Exec("INSERT INTO users (id, user_name, created_at) VALUES (?, ?, ?)",
		user.Id, user.UserName, now)
	if err != nil {
		return err
	}

	return nil
}

func (up *userPersistence) IsNewUser(userId int) bool {
	db := up.Conn

	_, err := db.Exec("SELECT * FROM users WHERE id = ? LIMIT 1", userId)
	if err == nil {
		return true
	}
	return false
}
