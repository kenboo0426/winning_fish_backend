package models

import (
	"database/sql"
	"log"
	"winning_fish_backend/config"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}
}
