package models

import (
	"database/sql"
	"fmt"
	"log"
	"winning_fish_backend/config"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
	tableNameQuiz = "quizzes"
	tableNameOption = "options"
	tableNameAnswer = "answers"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	createUserTable := fmt.Sprintf(`create table if not exists %s(
		id integer primary key autoincrement,
		uuid string not null unique,
		name string,
		emain string,
		password string,
		rating float,
		role ineter,
		created_at datetime
	)`, tableNameUser)
	Db.Exec(createUserTable)

	createQuizTable := fmt.Sprintf(`create table if not exists %s(
		id integer primary key autoincrement,
		image string,
		correct_id integer,
		correct_rate float,
		level integer,
		created_at datetime
	)`, tableNameQuiz)
	Db.Exec(createQuizTable)

	createOptionTable := fmt.Sprintf(`create table if not exists %s(
		id integer primary key autoincrement,
		name string,
		quiz_id integer
	)`, tableNameOption)
	Db.Exec(createOptionTable)

	createAnswerTable := fmt.Sprintf(`create table if not exists %s(
		id integer primary key autoincrement,
		user_id integer,
		quiz_id integer,
		correct boolean,
		answered_option_id integer
	)`, tableNameAnswer)
	Db.Exec(createAnswerTable)
}
