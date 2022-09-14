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

func (o *OnlineMatch) GetAskedQuizByOnlineMatch() (online_match_asked_quizzes []OnlineMatchAskedQuiz, err error) {
	cmd := `select id, quiz_id, online_match_id from online_match_asked_quizzes where online_match_id = ?`
	rows, err := Db.Query(cmd, o.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var askedQuiz OnlineMatchAskedQuiz
		err = rows.Scan(&askedQuiz.ID, &askedQuiz.QuizID, &askedQuiz.OnlineMatchID)
		if err != nil {
			log.Fatalln(err)
		}
		online_match_asked_quizzes = append(online_match_asked_quizzes, askedQuiz)
	}
	rows.Close()

	return online_match_asked_quizzes, err
}
