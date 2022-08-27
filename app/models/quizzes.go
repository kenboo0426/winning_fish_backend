package models

import (
	"fmt"
	"log"
	"time"
)

type Quiz struct {
	ID          int       `jsonapi:"id"`
	Image       string    `jsonapi:"image"`
	CorrectID   int       `jsonapi:"correct_id"`
	CorrectRate float32   `jsonapi:"correct_rate"`
	Level       int       `jsonapi:"level"`
	CreatedAt   time.Time `jsonapi:"created_at"`
}

func GetQuizzes() (quizzes []Quiz, err error) {
	fetchQuizzes := `select * from quizzes`
	rows, err := Db.Query(fetchQuizzes)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		quiz := &Quiz{}
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
		quizzes = append(quizzes, *quiz)
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

func DeleteQuiz(id int) (quiz Quiz, err error) {
	delete := "delete from quizzes where id = ?"
	
	_, err = Db.Exec(delete, id)
	if err != nil {
		fmt.Println(err)
	}
	return quiz, err
}
