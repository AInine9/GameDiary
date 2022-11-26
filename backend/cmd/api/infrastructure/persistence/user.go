package persistence

import (
	"backend/cmd/api/domain/model"
	"backend/cmd/api/domain/repository"
	"github.com/jinzhu/gorm"
)

type userPersistence struct {
	Conn *gorm.DB
}

func NewUserPersistence(conn *gorm.DB) repository.UserRepository {
	return &userPersistence{Conn: conn}
}

func (up *userPersistence) Search(name string) ([]*model.User, error) {
	var user []*model.User

	if err := up.Conn.Take(&user).Error; err != nil {
		return nil, err
	}

	db := up.Conn.Find(&user)

	if name != "" {
		db = db.Where("name = ?", name).Find(&user)
	}

	return user, nil
}
