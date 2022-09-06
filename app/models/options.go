package models

import (
	"log"
	"time"
)

type Option struct {
	ID        int
	Name      string
	QuizID    int
	CreatedAt time.Time
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
