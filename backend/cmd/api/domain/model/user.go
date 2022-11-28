package model

import "time"

type User struct {
	Id        int       `json:"id" db:"id"`
	UserName  string    `json:"user_name" db:"user_name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
