package models

import (
	"fmt"
	"log"
	"time"
)

type Quiz struct {
	ID          int       `json:"id"`
	Image       string    `json:"image"`
	CorrectID   int       `json:"correct_id"`
	CorrectRate *float32  `json:"correct_rate"`
	Level       int       `json:"level"`
	CreatedAt   time.Time `json:"created_at"`
	Options     []Option  `json:"options"`
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
		quiz.Options, _ = quiz.GetOptionsByQuiz()
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

func (q *Quiz) CreateQuiz() (id int64, err error) {
	createQuiz := `insert into quizzes (
		image,
		correct_id,
		level,
		created_at
	) values(?, ?, ?, ?)`

	result, err := Db.Exec(createQuiz, q.Image, q.CorrectRate, q.Level, time.Now())
	id, _ = result.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}

	return id, err
}

func (q *Quiz) UpdateQuiz() (err error) {
	updateQuiz, err := Db.Prepare(`update quizzes set image = ?,
	                                                  correct_id = ?,
																										level = ?, 
																										created_at = ? 
																										where id = ?`)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = updateQuiz.Exec(q.Image, q.CorrectID, q.Level, q.CreatedAt, q.ID)

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
