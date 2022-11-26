package config

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"time"
)

var (
	db  sqlx.DB
	err error
)

func Connect() *sqlx.DB {
	cfg := mysql.NewConfig()

	cfg.ParseTime = true
	cfg.Loc = getLocation()
	if host := getHost(); host != "" {
		cfg.Net = "tcp"
		cfg.Addr = host
	}
	cfg.User = getUser()
	cfg.Passwd = getPassword()
	cfg.DBName = getDatabase()

	db, err := sqlx.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	return db
}

func getHost() string {
	v := os.Getenv("MYSQL_HOST")
	if v == "" {
		log.Fatal()
	}
	return v
}

// Read MySQL user
func getUser() string {
	v := os.Getenv("MYSQL_USER")
	if v == "" {
		log.Fatal()
	}
	return v
}

// Read MySQL password
func getPassword() string {
	v := os.Getenv("MYSQL_PASSWORD")
	if v == "" {
		log.Fatal(err)
	}
	return v
}

// Read MySQL database name
func getDatabase() string {
	v := os.Getenv("MYSQL_DATABASE")
	if v == "" {
		log.Fatal(err)
	}
	return v
}

func getLocation() *time.Location {
	tz := os.Getenv("MYSQL_TZ")
	if tz == "" {
		return time.FixedZone("Asia/Tokyo", 9*60*60)
	}
	loc, err := time.LoadLocation(tz)
	if err != nil {
		log.Fatal(fmt.Errorf("invalid timezone %+v", tz))
	}
	return loc
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
