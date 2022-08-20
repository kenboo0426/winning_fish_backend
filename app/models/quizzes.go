package models

import (
	"log"
	"time"
)

type Quiz struct {
	ID          int
	Image       string
	CorrectID   int
	CorrectRate float32
	Level       int
	CreatedAt   time.Time
}

func GetQuizzes() (quizzes []Quiz, err error) {
	fetchQuizzes := `select * from quizzes`
	rows, err := Db.Query(fetchQuizzes)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var quiz Quiz
		err = rows.Scan(
			&quiz.ID,
			&quiz.Image,
			&quiz.CorrectID,
			&quiz.CorrectRate,
			&quiz.Level,
			&quiz.CreatedAt,
		)
		if err != nil {
			log.Fatalln(err)
		}
		quizzes = append(quizzes, quiz)
	}
	rows.Close()
	return quizzes, err
}

func GetQuiz(id int) (quiz Quiz, err error) {
	fetchquiz := `select * from quizzes where id = ?`
	quiz = Quiz{}
	err = Db.QueryRow(fetchquiz, id).Scan(
		&quiz.ID,
		&quiz.Image,
		&quiz.CorrectID,
		&quiz.CorrectRate,
		&quiz.Level,
		&quiz.CreatedAt,
	)

	return quiz, err
}

func (q *Quiz) CreateQuiz() (err error) {
	createQuiz := `insert into quizzes (
		image,
		correct_id,
		correct_rate,
		level,
		created_at
	) values(?, ?, ?, ?, ?)`

	_, err = Db.Exec(createQuiz, q.Image, q.CorrectID, q.CorrectRate, q.Level, time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	return err
}
