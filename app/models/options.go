package models

import (
	"log"
)

type Option struct {
	ID     int
	Name   string
	QuizID int
}

func (q *Quiz) GetOptionsByQuiz() (options []Option, err error) {
	cmd := `select id, name, quiz_id from options where quiz_id = ?`
	rows, err := Db.Query(cmd, q.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var option Option
		err = rows.Scan(&option.ID, &option.Name, &option.QuizID)
		if err != nil {
			log.Fatalln(err)
		}
		options = append(options, option)
	}
	rows.Close()

	return options, err
}

func (q *Quiz) CreateOptions() (correct_option_id int64, err error) {
	cmd := `insert into options (
		name,
		quiz_id
		) values(?, ?)`
	for index, option := range q.Options {
		result, _ := Db.Exec(cmd, option.Name, q.ID)
		if index == 0 {
			correct_option_id, err = result.LastInsertId()
		}
	}

	return correct_option_id, err
}
