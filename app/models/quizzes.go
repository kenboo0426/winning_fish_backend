package models

import (
	"time"
)

type Quiz struct {
	ID          int         `json:"id"`
	CorrectID   int         `json:"correct_id"`
	CorrectRate *float64    `json:"correct_rate"`
	Level       string      `json:"level"`
	CreatedAt   time.Time   `json:"created_at"`
	Options     []Option    `json:"options"`
	QuizImages  []QuizImage `json:"quiz_images"`
}

func GetQuizzes() (quizzes []Quiz, err error) {
	fetchQuizzes := `select * from quizzes`
	rows, err := Db.Query(fetchQuizzes)
	if err != nil {
		return quizzes, err
	}
	for rows.Next() {
		quiz := &Quiz{}
		err = rows.Scan(
			&quiz.ID,
			&quiz.CorrectID,
			&quiz.CorrectRate,
			&quiz.Level,
			&quiz.CreatedAt,
		)
		quiz.Options, _ = quiz.GetOptionsByQuiz()
		if err != nil {
			return quizzes, err
		}
		quizzes = append(quizzes, *quiz)
	}
	rows.Close()
	return quizzes, nil
}

func GetQuizzesByRandomAndLimitFive() (quizzes []Quiz, err error) {
	cmd := `select * from quizzes order by random() limit 5`
	rows, err := Db.Query(cmd)
	if err != nil {
		return quizzes, err
	}

	for rows.Next() {
		quiz := &Quiz{}
		err = rows.Scan(
			&quiz.ID,
			&quiz.CorrectID,
			&quiz.CorrectRate,
			&quiz.Level,
			&quiz.CreatedAt,
		)
		if err != nil {
			return quizzes, err
		}
		quizzes = append(quizzes, *quiz)
	}

	for len(quizzes) < 5 {
		quizzes = append(quizzes, quizzes...)
	}
	quizzes = quizzes[:5]

	rows.Close()
	return quizzes, nil
}

func GetQuiz(id int) (quiz Quiz, err error) {
	fetchquiz := `select * from quizzes where id = ?`
	quiz = Quiz{}
	err = Db.QueryRow(fetchquiz, id).Scan(
		&quiz.ID,
		&quiz.CorrectID,
		&quiz.CorrectRate,
		&quiz.Level,
		&quiz.CreatedAt,
	)
	if err != nil {
		return quiz, err
	}
	quiz.Options, err = quiz.GetOptionsByQuiz()
	if err != nil {
		return quiz, err
	}
	quiz.QuizImages, err = quiz.GetQuizImagesByQuiz()
	return quiz, err
}

func (q *Quiz) CreateQuiz() (id int64, err error) {
	createQuiz := `insert into quizzes (
		level,
		created_at
	) values(?, ?)`

	result, err := Db.Exec(createQuiz, q.Level, time.Now())
	if err != nil {
		return id, err
	}
	id, err = result.LastInsertId()
	return id, err
}

func (q *Quiz) UpdateQuiz() (err error) {
	updateQuiz, err := Db.Prepare(`update quizzes set correct_id = ?,
																										level = ?, 
																										created_at = ? 
																										where id = ?`)
	if err != nil {
		return err
	}

	_, err = updateQuiz.Exec(q.CorrectID, q.Level, q.CreatedAt, q.ID)
	return err
}

func DeleteQuiz(id int) (quiz Quiz, err error) {
	delete := "delete from quizzes where id = ?"

	_, err = Db.Exec(delete, id)
	return quiz, err
}
