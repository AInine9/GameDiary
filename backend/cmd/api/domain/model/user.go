package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	Id        uint
	name      string
	CreatedAt time.Time
}
