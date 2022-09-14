package models

import (
	"fmt"
	"log"
	"time"
)

type Answer struct {
	ID               int       `json:"id"`
	UserID           int       `json:"user_id"`
	QuizID           int       `json:"quiz_id"`
	Correct          bool      `json:"correct"`
	AnsweredOptionID int       `json:"answered_option_id"`
	OnlineMatchID    int       `json:"online_match_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (a *Answer) CreateAnswer() (err error) {
	cmd := `insert into answers (
		user_id,
		quiz_id,
		correct,
		answered_option_id,
		online_match_id,
		created_at,
		updated_at
	) values(?, ?, ?, ?, ?, ?, ?)`

	result, err := Db.Exec(cmd, a.UserID, a.QuizID, a.Correct, a.AnsweredOptionID, a.OnlineMatchID, time.Now(), time.Now())
	fmt.Println(err, a)
	id, _ := result.LastInsertId()
	a.ID = int(id)

	if err != nil {
		log.Fatalln()
	}

	return err
}
