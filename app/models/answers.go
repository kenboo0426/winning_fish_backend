package models

import (
	"log"
	"time"
)

type Answer struct {
	ID               int       `json:"id"`
	UserID           int       `json:"user_id"`
	QuizID           int       `json:"quiz_id"`
	Correct          bool      `json:"correct"`
	AnsweredOptionID int       `json:"answered_option_id"`
	RemainedTime     float32   `json:"remained_time"`
	OnlineMatchID    int       `json:"online_match_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func CalculateTotalRemainedTime(online_match_id int, user_id int) (total_time float32, err error) {
	cmd := `select sum(remained_time) as total_time from answers where online_match_id = ? and user_id = ?`
	rows, err := Db.Query(cmd, online_match_id, user_id)
	if err != nil {
		log.Fatalln(err)
	}

	var total_remained_time float64
	for rows.Next() {
		rows.Scan(&total_remained_time)
	}

	return float32(total_remained_time), err
}

func (a *Answer) CreateAnswer() (err error) {
	cmd := `insert into answers (
		user_id,
		quiz_id,
		correct,
		answered_option_id,
		remained_time,
		online_match_id,
		created_at,
		updated_at
	) values(?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := Db.Exec(cmd, a.UserID, a.QuizID, a.Correct, a.AnsweredOptionID, a.RemainedTime, a.OnlineMatchID, time.Now(), time.Now())
	id, _ := result.LastInsertId()
	a.ID = int(id)

	if err != nil {
		log.Fatalln()
	}

	return err
}
