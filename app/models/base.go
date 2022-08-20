package models

import (
	"database/sql"
	"fmt"
	"log"
	"winning_fish_backend/config"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
)

func init() {
	Db, err := sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	createUserTable := fmt.Sprintf(`create table if not exists %s(
		id integer primary key autoincreament,
		uuid string not null unique,
		name string,
		emain string,
		password string,
		rating float,
		role ineter,
		created_at datetime
	)`, tableNameUser)
	Db.Exec(createUserTable)
}
