package models

import (
	"log"
	"time"
)

type QuizImage struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	QuizID     int       `json:"quiz_id"`
	ProgressID int       `json:"progress_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (q *Quiz) CreateQuizImages() (err error) {
	cmd := `insert into quiz_images (
		name,
		quiz_id,
		progress_id,
		created_at,
		updated_at
	) values (?, ?, ?, ?, ?)`
	for _, quizImage := range q.QuizImages {
		_, err = Db.Exec(cmd, quizImage.Name, q.ID, quizImage.ProgressID, time.Now(), time.Now())
	}

	return err
}

func (q *Quiz) GetQuizImagesByQuiz() (images []QuizImage, err error) {
	cmd := `select id, name, quiz_id, progress_id from quiz_images where quiz_id = ? order by id desc`
	rows, err := Db.Query(cmd, q.ID)
	for rows.Next() {
		i := &QuizImage{}
		if err := rows.Scan(&i.ID, &i.Name, &i.QuizID, &i.ProgressID); err != nil {
			log.Fatalln(err)
		}
		images = append(images, *i)
	}

	return images, err
}
