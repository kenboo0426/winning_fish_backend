package models

import (
	"log"
	"time"
)

type OnlineMatchAskedQuiz struct {
	ID            int       `json:"id"`
	QuizID        int       `json:"quiz_id"`
	OnlineMatchID int       `json:"online_match_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (o *OnlineMatch) RegisterQuiz() {
	quizzes, err := GetQuizzesByRandomAndLimitFive()
	cmd := `insert into online_match_asked_quizzes (
		quiz_id,
		online_match_id,
		created_at,
		updated_at
	) values (?, ?, ?, ?)`

	if err != nil {
		log.Fatalln(err)
	}
	for _, quiz := range quizzes {
		_, err = Db.Exec(cmd, quiz.ID, o.ID, time.Now(), time.Now())
	}

	if err != nil {
		log.Fatalln(err)
	}

}
